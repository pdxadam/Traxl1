package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	dbTraxl "traxl/gen"
	"traxl/pkg"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	_ "github.com/lib/pq"
)

var (
	//go:embed static
	staticEmbed embed.FS

	//go:embed css/*
	cssEmbed embed.FS

	//go:embed tmpl/*.html
	tmplEmbed embed.FS

	dbQuery *dbTraxl.Queries

	store = sessions.NewCookieStore([]byte("forDemo"))
)

type staticHandler struct {
	staticPath string
	indexPage  string
}
type myResult struct {
	Message    string
	SecondPart string
}

func (h staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = filepath.Join(h.staticPath, path)
	_, err = os.Stat(path)
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)

}
func renderFiles(tmpl string, w http.ResponseWriter, d interface{}) {
	log.Println("Rendering " + tmpl)
	t, err := template.ParseFS(tmplEmbed, fmt.Sprintf("tmpl/%s.html", tmpl))
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(w, d); err != nil {
		log.Fatal(err)
	}
}

func securityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if sessionValid(w, r) {
			//login path will be let through.
			if r.URL.Path == "/login" {
				next.ServeHTTP(w, r)
				return
			}
		}

		if hasBeenAuthenticated(w, r) {
			next.ServeHTTP(w, r)
			return
		}
		//otherwise redirect to login
		storeAuthenticated(w, r, false)
		http.Redirect(w, r, "/login", 307)

	})
}
func sessionValid(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session_token")
	return !session.IsNew

}
func authenticationHandler(w http.ResponseWriter, r *http.Request) {
	result := "Login "
	r.ParseForm()
	log.Println(r.FormValue("username"))
	log.Println(r.FormValue("password"))
	if validateUser(r.FormValue("username"), r.FormValue("password")) {
		//TODO: do the store authenticated part
		result = result + "successful"
	} else {
		result = result + "unsuccesful"
	}
	log.Println(result)
	renderFiles("msg", w, result)

}
func hasBeenAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session_token")
	a, _ := session.Values["authenticated"]

	if a == nil {
		return false
	}
	return a.(bool)

}
func storeAuthenticated(w http.ResponseWriter, r *http.Request, v bool) {
	session, _ := store.Get(r, "session_token")
	session.Values["authenticated"] = v
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func validateUser(username string, password string) bool {

	ctx := context.Background()
	u, _ := dbQuery.GetUserByName(ctx, username)
	if u.Username != username {
		return false
	}
	return pkg.CheckPasswordHash(password, u.Passwordhash)
}

func main() {
	// port := "9004"

	// if value, exists := os.LookupEnv("SERVER_PORT"); exists {
	// 	port = value
	// }
	// log.Println(port)
	/*---- Unused Code right now
	connecting to the database
			------------------------*/
	// dbURI := fmt.Sprintf("%s?sslmode=disable",
	// 	GetAsString("PSQLURL", "Fail"))
	//
	// db, err := sql.Open("postgres", dbURI)
	// if err != nil {
	// 	panic(err)
	// }
	// if err := db.Ping(); err != nil {
	// 	log.Fatalln("Error from database ping:", err)
	// }
	// st := dbTraxl.New(db)
	// ctx := context.Background()
	/*----- Example
	     Creates a new user, create a topic, create an instance, list users
		 ------------------------------- */
	// newUser, err := st.CreateUsers(ctx, dbTraxl.CreateUsersParams{
	// 	Username:     "testuser",
	// 	Passwordhash: "hash",
	// 	Name:         "test",
	// })
	// if err != nil {
	// 	log.Fatalln("Error creating user :", err)
	// }
	// eid, err := st.InsertTopic(ctx, dbTraxl.InsertTopicParams{
	// 	Topicname: "testTopic",
	// 	Fkuser:    newUser.Pkuser,
	// })
	// if err != nil {
	// 	log.Fatalln("Error inserting Topic", err)

	// }
	// _, err = st.InsertInstance(ctx, dbTraxl.InsertInstanceParams{
	// 	StartDate: time.Now(),
	// 	Fktopic:   eid,
	// 	Fkuser:    newUser.Pkuser,
	// })
	// if err != nil {
	// 	log.Fatalln("Error inserting Instance: ", err)
	// }
	// log.Println("All done.")
	// u, err := st.ListUsers(ctx)
	// for _, usr := range u {
	// 	fmt.Println(fmt.Sprintf("Name : %s, ID : %d", usr.Name, usr.Pkuser))

	// }
	initDatabase()
	// makeTestUser()
	//router/server setup
	router := mux.NewRouter()

	//post handler for /login
	router.HandleFunc("/login", authenticationHandler).Methods("POST")

	//embed handler for /css path
	cssContentStatic, _ := fs.Sub(cssEmbed, "css")
	css := http.FileServer(http.FS(cssContentStatic))
	router.PathPrefix("/css").Handler(http.StripPrefix("/css", css))

	//embed handler for /app path (still need to create app section)
	contentStatic, _ := fs.Sub(staticEmbed, "static")
	static := http.FileServer(http.FS(contentStatic))
	router.PathPrefix("/app").Handler(securityMiddleware(http.StripPrefix("/app", static)))

	//add /login path
	router.PathPrefix("/login").Handler(securityMiddleware(http.StripPrefix("/login", static)))

	//root will redirect to /app
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app", http.StatusPermanentRedirect)
	})

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3334",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//start the server
	log.Fatal(srv.ListenAndServe())
}
func initDatabase() {
	log.Println("Initializing Database")

	dbURI := fmt.Sprintf("%s?sslmode=disable",
		GetAsString("PSQLURL", "Fail"))
	log.Println(dbURI)
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalln("Error from database ping:", err)
	}
	dbQuery = dbTraxl.New(db)

}
func makeTestUser() {
	hashPass, _ := pkg.HashPassword("newPassword")
	userData := dbTraxl.CreateUsersParams{
		Username:     "newUser",
		Passwordhash: hashPass,
		Name:         "Developer",
	}
	ctx := context.Background()

	newUser, err := dbQuery.CreateUsers(ctx, userData)
	if err != nil {
		log.Println("error creating test user")
	} else {
		log.Println(newUser)
	}

}

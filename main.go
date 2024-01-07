package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	dbTraxl "traxl/gen"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var (
	dbQuery dbTraxl.Queries
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
func postHandler(w http.ResponseWriter, r *http.Request) {
	result := myResult{Message: "Success", SecondPart: "You Made it"}
	r.ParseForm()
	if validateUser(r.FormValue("username"), r.FormValue("password")) {
		result.Message = "The login was successful"
	} else {
		result.Message = "Unsuccesful"
		result.SecondPart = "Please Try again"
	}
	t, err := template.ParseFiles("static/tmpl/loginResult.html")
	if err != nil {
		fmt.Fprintf(w, "error processing")
		return
	}
	tpl := template.Must(t, err)
	tpl.Execute(w, result)

}
func validateUser(username string, password string) bool {
	ctx := context.Background()
	dbQuery.GetUserByName(ctx, username)
	return true
}

func main() {
	port := "9004"

	if value, exists := os.LookupEnv("SERVER_PORT"); exists {
		port = value
	}
	log.Println(port)
	/*---- Unused Code right now
	connecting to the database
			------------------------*/
	// dbURI := fmt.Sprintf("%s?sslmode=disable",
	// 	GetAsString("PSQLURL", "Fail"))

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

	router := mux.NewRouter()
	spa := staticHandler{staticPath: "static", indexPage: "index.html"}
	router.PathPrefix("/").Handler(spa)
	srv := &http.Server{
		Handler:      router,
		Addr:         ":3334",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

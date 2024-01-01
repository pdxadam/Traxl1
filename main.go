package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	dbTraxl "traxl/gen"

	_ "github.com/lib/pq"
)

func main() {
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
	st := dbTraxl.New(db)
	ctx := context.Background()
	newUser, err := st.CreateUsers(ctx, dbTraxl.CreateUsersParams{
		Username:     "testuser",
		Passwordhash: "hash",
		Name:         "test",
	})
	if err != nil {
		log.Fatalln("Error creating user :", err)
	}
	eid, err := st.InsertTopic(ctx, dbTraxl.InsertTopicParams{
		Topicname: "testTopic",
		Fkuser:    newUser.Pkuser,
	})
	if err != nil {
		log.Fatalln("Error inserting Topic", err)

	}
	_, err = st.InsertInstance(ctx, dbTraxl.InsertInstanceParams{
		StartDate: time.Now(),
		Fktopic:   eid,
		Fkuser:    newUser.Pkuser,
	})
	if err != nil {
		log.Fatalln("Error inserting Instance: ", err)
	}
	log.Println("All done.")
	u, err := st.ListUsers(ctx)
	for _, usr := range u {
		fmt.Println(fmt.Sprintf("Name : %s, ID : %d", usr.Name, usr.Pkuser))

	}

}

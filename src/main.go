package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Event struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Startdate   string `db:"startdate"`
	Enddate     string `db:"enddate"`
}

func main() {
	db, err := sqlx.Connect("postgres", "user=admin dbname=mydb sslmode=disable password=admin host=postgres")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	event := Event{}

	rows, _ := db.Queryx("SELECT * FROM api.Events")

	for rows.Next() {
		err := rows.StructScan(&event)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%#v\n", event)
	}
}

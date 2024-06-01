package main

import (
	"log"

	"github.com/danielronalds/go-restful-api/src/db"
	"github.com/danielronalds/go-restful-api/src/resources"
)

func main() {
	pg := db.GetDatabase()
	defer pg.Close()

	// Test the connection to the database
	if err := pg.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	event := resources.Event{}

	rows, _ := pg.Queryx("SELECT * FROM api.Events")

	for rows.Next() {
		err := rows.StructScan(&event)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%#v\n", event)
	}
}

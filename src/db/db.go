package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


// Struct that represents the required info for connecting to the postgres db
type PostgresConfig struct {
	user     string
	password string
	dbName   string
	host     string
}

// Converts the configuration to the data source string for `sqlx`
func (c PostgresConfig) getDataSourceName() string {
	return fmt.Sprintf("user=%v dbname=%v sslmode=disable password=%v host=%v", c.user, c.dbName, c.password, c.host)
}

// Singleton variable of the database connections
var pg *sqlx.DB

// Gets the database connection
func GetDatabase() *sqlx.DB {
	// Implemented as a singleton
	if pg != nil {
		return pg
	}

	config := PostgresConfig{
		user:     "admin",
		password: "admin",
		dbName:   "mydb",
		host:     "192.168.0.2", // postgres container ip
	}

	newDB, err := sqlx.Connect("postgres", config.getDataSourceName())
	if err != nil {
		log.Fatalln(err)
	}

	pg = newDB

	return pg
}

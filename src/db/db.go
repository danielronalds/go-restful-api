package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// PostgresConfig Struct that represents the required info for connecting to the postgres db
type PostgresConfig struct {
	user     string
	password string
	dbName   string
	host     string
	timeout  int
}

// Converts the configuration to the data source string for `sqlx`
func (c PostgresConfig) getDataSourceName() string {
	return fmt.Sprintf("user=%v dbname=%v sslmode=disable connect_timeout=%v password=%v host=%v", c.user, c.dbName, c.timeout, c.password, c.host)
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
		host:     "172.18.0.2", // postgres container ip
		timeout:  5,
	}

	newDB, err := sqlx.Connect("postgres", config.getDataSourceName())
	if err != nil {
		log.Fatalln(err)
	}

	pg = newDB

	return pg
}

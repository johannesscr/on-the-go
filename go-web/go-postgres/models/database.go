package models

import (
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	// CREATE USER bond WITH PASSWORD 'james';
	// GRANT ALL PRIVILEGES ON DATABASE bookstore TO bond;
	var err error
	db, err = sql.Open("postgres", "postgres://bond:james@localhost:5432/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("You've connected to the database")
}
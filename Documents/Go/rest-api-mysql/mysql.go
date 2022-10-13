package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	// ("driverName", "username:password@tcp(localhost or 127.0.0.1:port)/databaseName")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_restapi")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

package database

import (
	"database/sql"
	"log"
)

var Database *sql.DB

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./DUCKY.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

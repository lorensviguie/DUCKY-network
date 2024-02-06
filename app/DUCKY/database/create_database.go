package database

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username     string
	Password     string
	Email        string
	SessionToken string
}

func Create() {
	//connect to database

	var db = ConnectDB()
	defer db.Close()
	// Create table User
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, key TEXT)")
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table User created successfully")
}

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
	log.Println("Table User created successfully")

	// Create table Room
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS rooms (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)")
	if err != nil {
		panic(err)
	}
	log.Println("Table Room created successfully")

	// Create table Room_Users
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS room_users (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER, room_id INTEGER, admin BOOLEAN, FOREIGN KEY(user_id) REFERENCES users(id), FOREIGN KEY(room_id) REFERENCES rooms(id))")
	if err != nil {
		panic(err)
	}
	log.Println("Table Room_Users created successfully")
}

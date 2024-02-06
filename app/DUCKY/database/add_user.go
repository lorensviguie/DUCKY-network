package database

import (
	logs "DUCKY/DUCKY/logs"
	"fmt"
)

func CreateUser(username string, key string) {
	db := ConnectDB()
	_, err := db.Exec("INSERT INTO users (username, key) VALUES (?, ?)", username, key)
	if err != nil {
		logs.LogToFile("db", err.Error())
		panic(err)
	}
	fmt.Println(username)
	logs.LogToFile("db", "Users "+username+" add to db with succes")
}

func CheckUser(username string) bool {
	db := ConnectDB()
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		panic(err)
	}
	if count > 0 {
		return true
	}
	return false
}

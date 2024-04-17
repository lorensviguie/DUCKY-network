package database

func Get_Room_Id(roomName string) int {
	db := ConnectDB()
	var id int
	err := db.QueryRow("SELECT (id) FROM rooms WHERE name = ?", roomName).Scan(&id)
	if err != nil {
		return 100
	}
	return id
}

func Get_Users_Id(roomName string) int {
	db := ConnectDB()
	var id int
	err := db.QueryRow("SELECT (id) FROM users WHERE username = ?", roomName).Scan(&id)
	if err != nil {
		return 100
	}
	return id
}

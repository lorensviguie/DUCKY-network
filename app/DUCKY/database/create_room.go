package database

import "database/sql"

func Create_Room(roomName string, name string) string {
	db := ConnectDB()
	erreur, check := Check_Room(roomName, db)
	if check != 0 {
		return erreur
	}
	// Insert new room
	_, err := db.Exec("INSERT INTO rooms (name) VALUES (?)", roomName)
	if err != nil {
		return (string(err.Error()))
	}
	userId := Get_Users_Id(name)
	roomId := Get_Room_Id(roomName)
	_, err = db.Exec("INSERT INTO room_users (user_id, room_id, admin) VALUES (?,?,?)", userId, roomId, true)
	if err != nil {
		return (string(err.Error()))
	}
	return ("Room created successfully")
}

func Check_Room(roomName string, db *sql.DB) (string, int) {
	// Check if room name already exists
	var roomCount int
	err := db.QueryRow("SELECT COUNT(*) FROM rooms WHERE name = ?", roomName).Scan(&roomCount)
	if err != nil {
		return (string(err.Error())), 2
	}

	if roomCount > 0 {
		return ("Room with the same name already exists."), 1

	}
	return "", 0
}

func Add_User_to_Room(RoomName, Username string) string {
	db := ConnectDB()
	userId := Get_Users_Id(Username)
	roomId := Get_Room_Id(RoomName)
	check := CheckUserRoom(roomId, userId, db)
	if !check {
		return "Error or Program IS BROKEN ONCE AGAIN FUCK"
	}
	_, err := db.Exec("INSERT INTO room_users (user_id, room_id, admin) VALUES (?,?,?)", userId, roomId, true)
	if err != nil {
		return (string(err.Error()))
	}
	return ("Uuser add to room with succes")
}

func CheckUserRoom(roomID, userID int, db *sql.DB)bool {
	// Check if room name already exists
	var roomCount int

	err := db.QueryRow("SELECT COUNT(*) FROM rooms_users WHERE user_id = ? AND room_id = ?", userID, roomID).Scan(&roomCount)
	if err != nil {
		return false
	}

	if roomCount > 0 {
		return false

	}
	return true
}

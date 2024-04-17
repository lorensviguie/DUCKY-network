package database

func GetUsersInRoom(roomName string) ([]string, error) {
	var usernames []string

	// Connect to the database
	db := ConnectDB()
	defer db.Close()

	// Query the usernames in the specified room
	rows, err := db.Query("SELECT users.username FROM users JOIN room_users ON users.id = room_users.user_id JOIN rooms ON rooms.id = room_users.room_id WHERE rooms.name = ?", roomName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and extract the usernames
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		usernames = append(usernames, username)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usernames, nil
}

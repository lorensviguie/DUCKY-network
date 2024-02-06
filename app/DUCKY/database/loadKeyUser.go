package database

func SelectKeyFromUser(username string) string {
	db := ConnectDB()

	var key string
	err := db.QueryRow("SELECT key FROM users WHERE username = ?", username).Scan(&key)
	if err != nil {
		return "Error"
	}
	return key
}

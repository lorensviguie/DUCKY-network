package bytesreader

import (
	db "DUCKY/DUCKY/database"
	"fmt"
	"strings"
)

func NewUser(lines []string) {
	username := lines[1]
	publicKey := strings.Join(lines[2:], "\n")
	fmt.Printf("Nom d'utilisateur reçu : %s\n", username)
	fmt.Printf("Clé publique reçue : %s\n", publicKey)
	if !db.CheckUser(username) {
		db.CreateUser(username, publicKey)
		fmt.Println("COIN")
	}
}

func StartAuthentificationUser(lines []string) string {
	username := lines[1]
	return username
}

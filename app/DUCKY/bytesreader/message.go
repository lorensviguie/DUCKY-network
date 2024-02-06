package bytesreader

import (
	send "DUCKY/DUCKY/sendMSG"
	"fmt"
	"net"
	"strings"
)

func MessageReader(conn net.Conn, reconstructedMessageSize int) {
	messageBuf := make([]byte, reconstructedMessageSize)
	_, err := conn.Read(messageBuf)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du message :", err)
	}
	fmt.Println("taille du message recu : ", reconstructedMessageSize)
	SplitMessage(string(messageBuf), conn)
}

func SplitMessage(messageBuff string, conn net.Conn) {
	lines := strings.Split(messageBuff, "\n")
	if lines[0] == "new user" {
		NewUser(lines)
	}
	fmt.Println(lines[0])
	if lines[0] == "askauthentification" {
		send.SendMessage(sendAuthRequest(lines), conn)
	}
	if lines[0] == "startauthentification" {
		send.SendMessage(CheckAuth(lines), conn)
	} else {
		fmt.Println(strings.Join(lines, " "))
	}
}

func GetRandomAuthByAuthID(authIDToFind string) ([]byte, string) {
	for i, auth := range StorageAuth {
		if auth.AuthID == authIDToFind {
			// Get the RandomAuth data
			randomAuth := auth.RandomAuth
			username := auth.Username

			// Remove the item from StorageAuth
			StorageAuth = append(StorageAuth[:i], StorageAuth[i+1:]...)

			// Return the RandomAuth data
			return randomAuth, username
		}
	}
	return nil, ""
}

func PrintAllLine(lines []string) {
	for i := 0; i < len(lines); i++ {
		fmt.Println([]byte(lines[i]))
	}
}

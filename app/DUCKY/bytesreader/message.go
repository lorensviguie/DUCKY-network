package bytesreader

import (
	security "DUCKY/DUCKY/security"
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
	switch lines[0] {
	case "new user":
		NewUser(lines)
	case "askkey":
		send.SendMessage([]byte(fmt.Sprintf("getkey\n%s", security.GetPublicKey())), conn)
	case "askauthentification":
		send.SendMessage(sendAuthRequest(lines), conn)
	case "startauthentification":
		send.SendMessage(CheckAuth(lines, conn), conn)
	case "tchat":
		send.SendToTchat([]byte(strings.Join(lines[3:], " ")), lines[1], conn)
	case "proveyouridentity":
		send.SendMessage(ProveIdentity(lines), conn)
	default:
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

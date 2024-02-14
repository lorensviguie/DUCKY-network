package bytesreader

import (
	send "DUCKY/client/sendMSG"
	serveur "DUCKY/client/serveurauth"
	store "DUCKY/client/structure"
	"bytes"
	"fmt"
	"net"
	"strings"
)

var log = false

func VarLog() bool {
	return log
}

func MessageReader(conn net.Conn, reconstructedMessageSize int) {
	messageBuf := make([]byte, reconstructedMessageSize)
	_, err := conn.Read(messageBuf)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du message :", err)
	}
	SplitMessage(string(messageBuf), conn)
}
func READERforserveurauth(conn net.Conn, reconstructedMessageSize int) {
	messageBuf := make([]byte, reconstructedMessageSize)
	_, err := conn.Read(messageBuf)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du message :", err)
	}
	lines := strings.Split(string(messageBuf), "\n")
	fmt.Println(lines[0])
	if lines[0] == "getkey" {
		err := serveur.WriteToFile(strings.Join(lines[1:], "\n"))
		if err != nil {
			fmt.Println("Erreur :", err)
		}
	}
	if lines[0] == "proveidentity" {
		data := strings.Join(lines[1:], "\n")
		if bytes.Equal(store.ServeurAUth, []byte(data)) {
			fmt.Println("--------------------\nSERVEUR AUTHENTIFIER\n--------------------")
			store.ServeurCheck = true
		} else {
			fmt.Println("ERREUR lors de l'authentification du serveur")
		}
	}
}

func SplitMessage(messageBuff string, conn net.Conn) {
	lines := strings.Split(messageBuff, "\n")
	if !log {
		fmt.Println(lines[0])
	}
	if lines[0] == "startauthentification" {
		send.SendMessage(MessageStartAuth(lines), conn)
	}
	if lines[0] == "01" {
		username := lines[2]
		alphaCheck := lines[4]
		addSession := store.Session{
			Username:  username,
			SessionID: []byte(alphaCheck),
		}
		store.Sessions = append(store.Sessions, addSession)
		log = true
		fmt.Println("--------------------------------------------------------------\n\nNow you are Connected To DUCKY NETWORK Enjoy your Session\n\n--------------------------------------------------------------")
		//Menu(conn)
	}
	if lines[0] == "tchat" {
		PrintAllLine(lines)
	} else {
		PrintAllLine(lines)
	}
}

func Menu(conn net.Conn) {
	for {
		var input string
		fmt.Scanln(&input)
		send.SendMessage([]byte(input), conn)
	}
}

func PrintAllLine(lines []string) {
	for i := 1; i < len(lines); i++ {
		fmt.Println((lines[i]))
	}
}

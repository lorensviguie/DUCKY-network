package bytesreader

import (
	decrypt "DUCKY/client/decrypt"
	send "DUCKY/client/sendMSG"
	store "DUCKY/client/structure"
	"fmt"
	"net"
	"strings"
)
var log = false

func MessageReader(conn net.Conn, reconstructedMessageSize int) {
	messageBuf := make([]byte, reconstructedMessageSize)
	_, err := conn.Read(messageBuf)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du message :", err)
	}
	SplitMessage(string(messageBuf), conn)
}

func SplitMessage(messageBuff string, conn net.Conn) {
	lines := strings.Split(messageBuff, "\n")
	if !log{
		fmt.Println(lines[0])
	}
	if lines[0] == "startauthentification" {
		alphaCheck := []byte(lines[1])
		messagebyte := []byte(strings.Join(lines[2:], "\n"))
		messagedecrypt, _ := decrypt.DecryptMessage(messagebyte, GetPrivateKey())
		//fmt.Println(messagedecrypt)
		message := []byte(fmt.Sprintf("startauthentification\n%s\n%s", alphaCheck, messagedecrypt))
		send.SendMessage(message, conn)
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
	if lines[0] == "tchat"{
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
	for i := 0; i < len(lines); i++ {
		fmt.Println((lines[i]))
	}
}

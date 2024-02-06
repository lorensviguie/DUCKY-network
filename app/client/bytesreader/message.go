package bytesreader

import (
	decrypt "DUCKY/client/decrypt"
	send "DUCKY/client/sendMSG"
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
	SplitMessage(string(messageBuf), conn)
}

func SplitMessage(messageBuff string, conn net.Conn) {
	lines := strings.Split(messageBuff, "\n")
	fmt.Println(lines[0])
	if lines[0] == "startauthentification" {
		alphaCheck := []byte(lines[1])
		messagebyte := []byte(strings.Join(lines[2:], "\n"))
		messagedecrypt, _ := decrypt.DecryptMessage(messagebyte, GetPrivateKey())
		fmt.Println(messagedecrypt)
		message := []byte(fmt.Sprintf("startauthentification\n%s\n%s", alphaCheck, messagedecrypt))
		send.SendMessage(message, conn)
	} else {
		PrintAllLine(lines)
	}
}

func PrintAllLine(lines []string) {
	for i := 0; i < len(lines); i++ {
		fmt.Println((lines[i]))
	}
}

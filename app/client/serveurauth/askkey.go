package serveurauth

import (
	send "DUCKY/client/sendMSG"
	"fmt"
	"net"
)

func AskServerKey(conn net.Conn) {
	message := []byte("askkey")
	fmt.Print("je veux une clé serveur")
	send.SendMessage(message, conn)
}

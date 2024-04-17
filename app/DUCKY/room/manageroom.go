package room

import (
	sendmsg "DUCKY/DUCKY/sendMSG"
	"DUCKY/DUCKY/structure"
	"net"
	"strings"
)

func ManageRoom(lines []string, conn net.Conn) {
	username, _ := structure.FindUsernameByConn(conn)
	messageByte := []byte(strings.Join(lines[1:], "\n"))
	// messagedecrypt, err := security.DecryptMessage(security.GetPrivateKey(), messageByte)
	data := strings.Split(string(messageByte), "\n")
	sendmsg.SendToRoom(data, username)
}

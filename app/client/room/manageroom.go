package room

import (
	send "DUCKY/client/sendMSG"
	store "DUCKY/client/structure"
	"fmt"
	"net"
	"strings"
)

func ManageRoom(input string, conn net.Conn) {
	header := "room"
	message := strings.Split(input, " ")
	roomName := (message[0])[1:]
	messageG := strings.Join(message[1:], " ")
	data := []byte(store.Sessions.Username + "\n")
	data = append(data, store.Sessions.SessionID...)
	data = append(data, "\n"...)
	data = append(data, roomName+"\n"...)
	data = append(data, messageG...)
	// cryptData, _ := decrypt.EncryptMessageWithPublic(security.GetServeurPublicKey(), string(data))
	toSend := []byte(fmt.Sprintf("%s\n%s", header, data))
	send.SendMessage([]byte(toSend), conn)
}

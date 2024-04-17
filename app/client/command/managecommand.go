package command

import (
	store "DUCKY/client/structure"
	send "DUCKY/client/sendMSG"
	"fmt"
	"net"
	"strings"
)

func ManageCommand(input string,conn net.Conn) {
	header := "command"
	test := strings.Split(input, " ")
	if test[0] == "/accept_invite" {
		toSend := []byte(fmt.Sprintf("%s\n%s\n%s\n%s", header, store.Sessions.Username, store.Sessions.SessionID, input))
		send.SendMessage([]byte(toSend), conn)
	}
	if strings.ToLower(test[0]) == "/help"{
		fmt.Println(store.HelpString)
	} else {
		toSend := []byte(fmt.Sprintf("%s\n%s\n%s\n%s", header, store.Sessions.Username, store.Sessions.SessionID, input))
		send.SendMessage([]byte(toSend), conn)
	}
}

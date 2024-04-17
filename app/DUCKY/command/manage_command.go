package command

import (
	data "DUCKY/DUCKY/database"
	"DUCKY/DUCKY/security"
	send "DUCKY/DUCKY/sendMSG"
	structure "DUCKY/DUCKY/structure"
	"fmt"
	"net"
	"strings"
)

func ManageCommand(list []string, Sender string, conn net.Conn) {
	command := strings.Split(list[0], " ")
	message := "Erreur lors de l'execution de la commande \nCOMMAND NOT FOUND"
	fmt.Println(command[0])
	switch command[0] {
	case "/create_room":
		message = data.Create_Room(command[1], command[2])
	case "/invite":
		room := command[1]
		user := command[2]
		fmt.Println(command[1] + " : " + command[2])
		_, checkerr := data.Check_Room(room, data.ConnectDB())
		if checkerr == 1 {
			message = Invite_To_Room(user, Sender, room)
		} else {
			message = "Le room a l'air de ne pas exister ou pb xd"
		}
	case "/accept_invite":
		message = Accept_invite(Sender, command[1])
	}
	fmt.Println(message)
	username, _ := structure.FindUsernameByConn(conn)
	message_encrypt, _ := security.EncryptMessageWithPublic(data.SelectKeyFromUser(username), (message))
	send.SendMessage([]byte("command\n"+string(message_encrypt)), conn)
}

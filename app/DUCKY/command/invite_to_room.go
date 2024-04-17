package command

import (
	data "DUCKY/DUCKY/database"
	"DUCKY/DUCKY/security"
	send "DUCKY/DUCKY/sendMSG"
	authStorage "DUCKY/DUCKY/structure"
)

func Invite_To_Room(Tosend, Sender, RoomName string) string {
	allSession := authStorage.Sessions
	for _, session := range allSession {
		if session.Username == Tosend {
			if !data.CheckUserRoom(data.Get_Room_Id(RoomName), data.Get_Users_Id(session.Username), data.ConnectDB()) {
				message := ("User is already in room OR Program is Broken Tell to the Fucking duck that dev this app that Swan is better")
				privatekey := data.SelectKeyFromUser(session.Username)
				cryptdata, _ := security.EncryptMessageWithPublic(privatekey, string(Sender+"\n"+string(message)))
				Tosend := append([]byte("tchat\n"), cryptdata...)
				send.SendMessage(Tosend, session.Conn)
			}
			message := (Sender + "Invite you to join Room : " + RoomName + " |/accept_invite duck")
			invite := authStorage.Invite{
				Username:  string(session.Username),
				Room_Name: RoomName,
			}
			privatekey := data.SelectKeyFromUser(session.Username)
			cryptdata, _ := security.EncryptMessageWithPublic(privatekey, string(Sender+"\n"+string(message)))
			Tosend := append([]byte("tchat\n"), cryptdata...)

			authStorage.Invite_List = append(authStorage.Invite_List, invite)
			send.SendMessage(Tosend, session.Conn)
			return "Invite send with succes"
		}
	}
	return "Invite wasn't sent. Perhaps the user is offline or there is an issue with the program. ¯\\_(°_o)_/¯"
}

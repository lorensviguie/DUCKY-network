package command

import (
	data "DUCKY/DUCKY/database"
	storage "DUCKY/DUCKY/structure"
	"fmt"
)

func Accept_invite(Sender, RoomName string) string {
	for i, session := range storage.Invite_List {
			fmt.Println(session.Username + " : "+ Sender + "\n" + session.Room_Name + " : "+ RoomName)
		if (session.Username == Sender) && (session.Room_Name == RoomName) {
			exec := data.Add_User_to_Room(RoomName, Sender)
			storage.Invite_List = append(storage.Invite_List[:i], storage.Invite_List[i+1:]...)
			return exec
		}
	}
	return "ERROR"
}

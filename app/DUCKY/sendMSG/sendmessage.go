package sendmsg

import (
	"DUCKY/DUCKY/database"
	db "DUCKY/DUCKY/database"
	"DUCKY/DUCKY/security"
	authStorage "DUCKY/DUCKY/structure"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func CompileMessageSize(message []byte) []byte {
	// Définir la taille du tableau en fonction de la taille du message
	sizeBytes := make([]byte, 2)

	// Écrire la taille dans le tableau de bytes en utilisant binary.BigEndian.PutUint16
	binary.BigEndian.PutUint16(sizeBytes, uint16(len(message)))

	return sizeBytes
}

func CompileHeaderSize(messageSize []byte) byte {
	headerSize := byte(len(messageSize))
	return headerSize
}

func SendMessage(message []byte, conn net.Conn) {
	messageSize := CompileMessageSize(message)
	headerSize := []byte{CompileHeaderSize(messageSize)}
	fmt.Println("msgsize", len(message))
	fmt.Println("hsize :", headerSize)
	data := append(append(headerSize, messageSize...), message...)
	if _, err := conn.Write(data); err != nil {
		conn.Close()
		fmt.Println("Erreur lors de l'envoi du message :", err)
		return
	}
	fmt.Println("Message send with succces to", conn.RemoteAddr())
}

func SendToTchat(message []byte, Sender string, connToExclude net.Conn) {
	allSession := authStorage.Sessions

	for _, session := range allSession {
		if session.Conn != connToExclude {
			fmt.Println("\n message send to ", session.Username)
			privatekey := db.SelectKeyFromUser(session.Username)
			cryptdata, _ := security.EncryptMessageWithPublic(privatekey, string(Sender+"\n"+string(message)))
			Tosend := append([]byte("tchat\n"), cryptdata...)
			SendMessage(Tosend, session.Conn)
		}
	}
}

func SendToUser(message []byte, Sender string, Tosend string) {
	allSession := authStorage.Sessions

	for _, session := range allSession {
		if session.Username == Tosend {
			fmt.Println("\n message send to ", session.Username)
			privatekey := db.SelectKeyFromUser(session.Username)
			cryptdata, _ := security.EncryptMessageWithPublic(privatekey, string(Sender+"\n"+string(message)))
			Tosend := append([]byte("tchat\n"), cryptdata...)
			SendMessage(Tosend, session.Conn)
		}
	}
}

func SendToRoom(data []string,sender string) {
	//username := data[0]
	//sessionID := data[1]
	roomName := data[3]
	message := strings.Join(data[4:], " ")
	userstosend, _ := database.GetUsersInRoom(roomName)
	allSession := authStorage.Sessions
	for _, session := range allSession {
		if contains(userstosend, session.Username) {
			fmt.Println("\n message send to ", session.Username)
			privatekey := db.SelectKeyFromUser(session.Username)
			cryptdata, _ := security.EncryptMessageWithPublic(privatekey, string(sender+"\n"+string(message)))
			Tosend := append([]byte("room\n"), cryptdata...)
			SendMessage(Tosend, session.Conn)
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

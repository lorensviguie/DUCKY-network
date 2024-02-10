package sendmsg

import (
	authStorage "DUCKY/DUCKY/structure"
	"encoding/binary"
	"fmt"
	"net"
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
	Tosend := append(append([]byte("tchat\n"), []byte("["+Sender+"]")...), message...)
	for _, session := range allSession {
		if session.Conn != connToExclude {
			fmt.Println("\n message send to ", session.Username)
			SendMessage(Tosend, session.Conn)
		}
	}
}

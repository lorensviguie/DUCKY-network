package main

import (
	affichage "DUCKY/DUCKY/affichage"
	br "DUCKY/DUCKY/bytesreader"
	db "DUCKY/DUCKY/database"
	logs "DUCKY/DUCKY/logs"
	sec "DUCKY/DUCKY/security"
	"fmt"
	"net"
)

type ClientInfo struct {
	IP   string
	Conn net.Conn
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	affichage.FormatAndDisplay(fmt.Sprintf("Nouvelle connexion établie depuis %s\n", conn.RemoteAddr()))
	logs.LogToFile("Connexion", fmt.Sprintf("New Connexion From %s\n", conn.RemoteAddr()))

	for {

		headerSize := br.ReadHeaderSize(conn)
		if headerSize != 0 {
			messagesize := br.ReadMessageSize(conn, headerSize)
			br.MessageReader(conn, messagesize)
		}
	}
}

func main() {
	db.Create()
	err := sec.GenerateKeyPair()
	if err != nil {
		fmt.Println("Erreur lors de la génération de la paire de clés:", err)
	}
	listener, err := net.Listen("tcp", ":666")
	if err != nil {
		fmt.Println("Erreur lors de l'écoute sur le port 666 :", err)
		return
	}
	fmt.Println("Serveur en attente de connexions sur le port 666...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation de la connexion :", err)
			continue
		}
		go handleConnection(conn)
	}
}

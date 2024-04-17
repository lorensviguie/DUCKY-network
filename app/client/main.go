package main

import (
	br "DUCKY/client/bytesreader"
	command "DUCKY/client/command"
	room "DUCKY/client/room"
	security "DUCKY/client/security"
	send "DUCKY/client/sendMSG"
	serveur "DUCKY/client/serveurauth"
	store "DUCKY/client/structure"
	user "DUCKY/client/user"
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func keysExist() bool {
	var privateKeyPath = filepath.Join(store.KeyPath, "private.pem")
	var publicKeyPath = filepath.Join(store.KeyPath, "public.pem")
	_, privateErr := os.Stat(privateKeyPath)
	_, publicErr := os.Stat(publicKeyPath)
	return !os.IsNotExist(privateErr) && !os.IsNotExist(publicErr)
}
func HaveServeurKey() bool {
	serveurKeyPath := filepath.Join(store.KeyPath, "serveurpublickey.pem")
	_, privateErr := os.Stat(serveurKeyPath)
	return !os.IsNotExist(privateErr)
}

func sendUserInfo(username string, publicKey string, conn net.Conn) {
	message := []byte(fmt.Sprintf("new user\n%s\n%s", username, publicKey))
	send.SendMessage(message, conn)

}

func askAuthentification(username string, conn net.Conn) {
	message := []byte(fmt.Sprintf("askauthentification\n%s", username))
	send.SendMessage(message, conn)

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		headerSize := br.ReadHeaderSize(conn)
		if !store.ServeurCheck {
			messagesize := br.ReadMessageSize(conn, headerSize)
			br.READERforserveurauth(conn, messagesize)
		} else {
			if headerSize != 0 {
				messagesize := br.ReadMessageSize(conn, headerSize)
				if !br.VarLog() {
					fmt.Println("\nYou receive a message from : ", conn.RemoteAddr())
					fmt.Println("taille du header recu: ", headerSize)
					fmt.Println("taille du message recu : ", messagesize)
				}
				br.MessageReader(conn, messagesize)
			}

		}
	}
}

func main() {
	if len(os.Args[1:]) > 0 {
		fmt.Println("You change key location to : ", os.Args[1])
		store.KeyPath = os.Args[1]
	}
	fmt.Print("Veuillez entrer votre nom d'utilisateur : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	serverAddr := "localhost:666"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Erreur lors de la connexion au serveur :", err)
		return
	}
	if !keysExist() {
		user.NewUser()
		sendUserInfo(username, security.GetPublicKey(), conn)
	}
	if !HaveServeurKey() {
		serveur.AskServerKey(conn)
	}
	security.AskServerAuthentification(conn)
	askAuthentification(username, conn)
	go handleConnection(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if input == "\n" {

		} else {
			input = strings.TrimSpace(input) // supprime les espaces blancs autour de la saisie
			header := "tchat"
			if strings.HasPrefix(input, "/") {
				command.ManageCommand(input, conn)
			}
			if strings.HasPrefix(input, "\\") {
				room.ManageRoom(input, conn)
			} else {
				toSend := []byte(fmt.Sprintf("%s\n%s\n%s\n%s", header, store.Sessions.Username, store.Sessions.SessionID, input))
				send.SendMessage([]byte(toSend), conn)
			}
		}
	}
}

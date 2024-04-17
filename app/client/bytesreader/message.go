package bytesreader

import (
	decrypt "DUCKY/client/decrypt"
	send "DUCKY/client/sendMSG"
	serveur "DUCKY/client/serveurauth"
	store "DUCKY/client/structure"
	"bytes"
	"fmt"
	"net"
	"strings"
)

var log = false

func VarLog() bool {
	return log
}

func MessageReader(conn net.Conn, reconstructedMessageSize int) {
	messageBuf := make([]byte, reconstructedMessageSize)
	_, err := conn.Read(messageBuf)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du message :", err)
	}
	SplitMessage(string(messageBuf), conn)
}

func READERforserveurauth(conn net.Conn, reconstructedMessageSize int) {
	messageBuf := make([]byte, reconstructedMessageSize)
	_, err := conn.Read(messageBuf)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du message :", err)
	}
	lines := strings.Split(string(messageBuf), "\n")
	fmt.Println(lines[0])
	if lines[0] == "getkey" {
		err := serveur.WriteToFile(strings.Join(lines[1:], "\n"))
		if err != nil {
			fmt.Println("Erreur :", err)
		}
	}
	if lines[0] == "proveidentity" {
		data := strings.Join(lines[1:], "\n")
		fmt.Println(store.ServeurAUth)
		if bytes.Equal(store.ServeurAUth, []byte(data)) {
			fmt.Println("--------------------\nSERVEUR AUTHENTIFIER\n--------------------")
			store.ServeurCheck = true
		} else {
			fmt.Println("ERREUR lors de l'authentification du serveur")
		}
	}
}

func SplitMessage(messageBuff string, conn net.Conn) {
	lines := strings.Split(messageBuff, "\n")
	if !log {
		fmt.Println(lines[0])
	}
	if lines[0] == "startauthentification" {
		send.SendMessage(MessageStartAuth(lines), conn)
	}
	if lines[0] == "01" {
		encryptedData := []byte(strings.Join(lines[5:], "\n"))
		privateKeyStr := GetPrivateKey()
		decryptedData, err := decrypt.DecryptMessageWithPrivate(privateKeyStr, encryptedData)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(removeLastCharacter(decryptedData))
		username := lines[2]
		alphaCheck := lines[4]
		addSession := store.Session{
			Username:  username,
			SessionID: []byte(alphaCheck),
			PublicKey: string(decryptedData),
		}
		store.Sessions = addSession
		log = true
		fmt.Println("--------------------------------------------------------------\n\nNow you are Connected To DUCKY NETWORK Enjoy your Session\n\n--------------------------------------------------------------")
		fmt.Println("\n you're connected has : " + lines[2])
		store.Authentifier = true
	}
	if lines[0] == "tchat" {
		PrintAllLine(lines)
	}
	if lines[0] == "command" {
		encryptedData := []byte(strings.Join(lines[1:], "\n"))
		privateKeyStr := GetPrivateKey()
		decryptedData, err := decrypt.DecryptMessageWithPrivate(privateKeyStr, encryptedData)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(decryptedData)
	}
	if lines[0] == "room"{
		encryptedData := []byte(strings.Join(lines[1:], "\n"))
		privateKeyStr := GetPrivateKey()
		decryptedData, _ := decrypt.DecryptMessageWithPrivate(privateKeyStr, encryptedData)
		messageFinal := strings.Split(decryptedData, "\n")
		for i := 0; i < len((messageFinal)); i++ {
			println(messageFinal[i])	
		}
	}
}

func Menu(conn net.Conn) {
	for {
		var input string
		fmt.Scanln(&input)
		send.SendMessage([]byte(input), conn)
	}
}

func PrintAllLine(lines []string) {
	messagebyte := []byte(strings.Join(lines[1:], "\n"))
	messagedecrypt, _ := decrypt.DecryptMessageWithPrivate(GetPrivateKey(), messagebyte)
	messageFinal := strings.Split(messagedecrypt, "\n")
	fmt.Println("[" + messageFinal[0] + "] " + messageFinal[1])
}

func removeLastCharacter(input string) string {
	if len(input) > 0 {
		return input[:len(input)-1]
	}
	return input
}

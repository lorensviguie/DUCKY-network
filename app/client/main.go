package main

import (
	br "DUCKY/client/bytesreader"
	send "DUCKY/client/sendMSG"
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func keysExist() bool {
	privateKeyPath := filepath.Join(".ssh", "private.pem")
	publicKeyPath := filepath.Join(".ssh", "public.pem")

	_, privateErr := os.Stat(privateKeyPath)
	_, publicErr := os.Stat(publicKeyPath)

	return !os.IsNotExist(privateErr) && !os.IsNotExist(publicErr)
}

func generateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func savePEMKey(filename string, key *rsa.PrivateKey) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	privBytes := x509.MarshalPKCS1PrivateKey(key)
	privBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes}

	return pem.Encode(file, privBlock)
}

func savePEMKeyPublic(filename string, pubkey *rsa.PublicKey) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	pubBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return err
	}
	pubBlock := &pem.Block{Type: "RSA PUBLIC KEY", Bytes: pubBytes}

	return pem.Encode(file, pubBlock)
}

func sendUserInfo(username string, publicKey string, conn net.Conn) {
	message := []byte(fmt.Sprintf("new user\n%s\n%s", username, publicKey))
	send.SendMessage(message, conn)

}

func askAuthentification(username string, conn net.Conn) {
	message := []byte(fmt.Sprintf("askauthentification\n%s", username))
	send.SendMessage(message, conn)

}

func getPublicKey() string {
	publicKeyPath := filepath.Join(".ssh", "public.pem")
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la clé publique :", err)
		return "err"
	}
	return string(publicKeyBytes)
}

func NewUser() {
	privateKey, publicKey, err := generateKeyPair(2048)
	if err != nil {
		fmt.Println("Erreur lors de la génération de la paire de clés:", err)
		return
	}

	err = savePEMKey(".ssh/private.pem", privateKey)
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde de la clé privée:", err)
		return
	}

	err = savePEMKeyPublic(".ssh/public.pem", publicKey)
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde de la clé publique:", err)
		return
	}

	fmt.Println("Paire de clés générée et sauvegardée avec succès.")
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {

		headerSize := br.ReadHeaderSize(conn)
		if headerSize != 0 {
			fmt.Println("\nYou receive a message from : ", conn.RemoteAddr())
			fmt.Println("taille du header recu: ", headerSize)
			messagesize := br.ReadMessageSize(conn, headerSize)
			fmt.Println("taille du message recu : ", messagesize)
			br.MessageReader(conn, messagesize)
		}
	}
}

func main() {
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
		NewUser()
		sendUserInfo(username, getPublicKey(), conn)
	} else {
		askAuthentification(username, conn)
	}
	go handleConnection(conn)

	// Ici, vous pouvez effectuer d'autres opérations dans votre programme si nécessaire

	// Pour éviter que le programme principal ne se termine immédiatement
	// Vous pouvez attendre une entrée utilisateur ou utiliser un autre mécanisme pour laisser le programme en cours d'exécution
	var input string
	fmt.Scanln(&input)
	defer conn.Close()
}

package security

import (
	send "DUCKY/client/sendMSG"
	auth "DUCKY/client/structure"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
)

func AskServerAuthentification(conn net.Conn) {
	serveurkey := GetServeurPublicKey()
	ciphertext, randomdata, err := encrypt(serveurkey)
	if err != nil {
		fmt.Println(err)
	}
	auth.ServeurAUth = randomdata
	send.SendMessage(append([]byte("proveyouridentity\n"), ciphertext...), conn)
}

func encrypt(publicKeyStr string) ([]byte, []byte, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, nil, fmt.Errorf("erreur lors du décodage de la clé publique")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("erreur lors du parsing de la clé publique : %v", err)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, nil, fmt.Errorf("la clé n'est pas une clé rsa valide")
	}
	randomData := make([]byte, 32)
	_, err = rand.Read(randomData)
	if err != nil {
		return nil, nil, fmt.Errorf("erreur lors de la génération de données aléatoires : %v", err)
	}
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, randomData)
	if err != nil {
		return nil, nil, fmt.Errorf("erreur lors du chiffrement : %v", err)
	}
	fmt.Println(ciphertext,randomData)
	return ciphertext, randomData, nil
}

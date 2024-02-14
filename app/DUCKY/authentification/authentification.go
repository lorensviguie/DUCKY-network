package authentification

import (
	db "DUCKY/DUCKY/database"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
)

func encryptAndGenerateID(publicKeyStr string) ([]byte, []byte, string, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, nil, "", fmt.Errorf("erreur lors du décodage de la clé publique")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, nil, "", fmt.Errorf("erreur lors du parsing de la clé publique : %v", err)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, nil, "", fmt.Errorf("la clé n'est pas une clé rsa valide")
	}
	randomData := make([]byte, 32)
	_, err = rand.Read(randomData)
	if err != nil {
		return nil, nil, "", fmt.Errorf("erreur lors de la génération de données aléatoires : %v", err)
	}
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, randomData)
	if err != nil {
		return nil, nil, "", fmt.Errorf("erreur lors du chiffrement : %v", err)
	}

	id, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(10), big.NewInt(10), nil))
	if err != nil {
		return nil, nil, "", fmt.Errorf("erreur lors de la génération de l'identifiant unique : %v", err)
	}
	alphacheck := string(id.Int64())
	return ciphertext, randomData, alphacheck, nil
}

func StartAuthentificationUser(username string) ([]byte, []byte, string) {
	publickey := db.SelectKeyFromUser(username)
	if publickey == "Error" {
		return []byte{110}, []byte{0}, "no"
	}
	cryptText, uncrypttext, Alphacheck, err := encryptAndGenerateID(publickey)
	if err != nil {
		fmt.Println(err)
	}
	return cryptText, uncrypttext, Alphacheck
}

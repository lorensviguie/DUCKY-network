package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func EncryptMessageWithPublic(publicKeyStr string, message string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("erreur lors du décodage de la clé publique")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du parsing de la clé publique : %v", err)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("la clé n'est pas une clé rsa valide")
	}
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(message))
	if err != nil {
		return nil, fmt.Errorf("erreur lors du chiffrement : %v", err)
	}
	return ciphertext, nil
}

func EncryptMessageWithPrivate(privateKeyStr string, message string) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("erreur lors du décodage de la clé privée")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du parsing de la clé privée : %v", err)
	}

	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, []byte(message))
	if err != nil {
		return nil, fmt.Errorf("erreur lors du chiffrement : %v", err)
	}
	return ciphertext, nil
}

package decrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func DecryptMessage(ciphertext []byte, privateKeyStr string) ([]byte, error) {
	// Parse la clé privée au format PEM
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("Erreur lors du décodage de la clé privée")
	}

	// Parse la clé privée au format RSA
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors du parsing de la clé privée : %v", err)
	}

	// Déchiffre le message avec la clé privée RSA
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors du déchiffrement : %v", err)
	}

	return plaintext, nil
}

func DecryptMessageWithPrivate(privateKeyStr string, ciphertext []byte) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return "", fmt.Errorf("error decoding private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("error parsing private key: %v", err)
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return "", fmt.Errorf("error decrypting: %v", err)
	}

	return string(plaintext), nil
}

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

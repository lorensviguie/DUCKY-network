package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func DecryptMessage(privateKeyStr string, ciphertext []byte) ([]byte, error) {
	// Parse la clé privée au format PEM
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("Erreur lors du décodage de la clé privée")
	}

	// Parse la clé privée au format PKCS#1
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

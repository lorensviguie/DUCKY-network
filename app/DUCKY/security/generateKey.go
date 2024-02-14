package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

// generateKeyPair génère une paire de clés PEM et les stocke dans le répertoire spécifié.
func GenerateKeyPair() error {
	// Crée le répertoire s'il n'existe pas
	var directory = "./DUCKY/.ssh"
	if err := os.MkdirAll(directory, 0700); err != nil {
		return err
	}

	// Génère une paire de clés RSA
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Encode la clé privée en PEM
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes}

	// Écrit la clé privée dans un fichier
	privateKeyPath := filepath.Join(directory, "private_key.pem")
	privateKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()
	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		return err
	}

	fmt.Printf("Clé privée générée: %s\n", privateKeyPath)

	// Encode la clé publique en PEM
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	publicKeyPEM := &pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBytes}

	// Écrit la clé publique dans un fichier
	publicKeyPath := filepath.Join(directory, "public_key.pem")
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		return err
	}
	defer publicKeyFile.Close()
	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		return err
	}

	fmt.Printf("Clé publique générée: %s\n", publicKeyPath)

	return nil
}

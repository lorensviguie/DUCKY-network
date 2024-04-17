package security

import (
	"DUCKY/client/security"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateKeyRSA(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func PrivateKeyToString(privateKey *rsa.PrivateKey) string {
	// Encode the private key to PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Convert the PEM-encoded private key to a string
	privateKeyStr := string(pem.EncodeToMemory(privateKeyPEM))

	return privateKeyStr
}

func PublicKeyToString(publicKey *rsa.PublicKey) string {
	// Encode the public key to PEM format
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	// Convert the PEM-encoded public key to a string
	publicKeyStr := string(pem.EncodeToMemory(publicKeyPEM))

	return publicKeyStr
}

func SavePEMKey(filename string, key *rsa.PrivateKey) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	privBytes := x509.MarshalPKCS1PrivateKey(key)
	privBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes}

	return pem.Encode(file, privBlock)
}

func SavePEMKeyPublic(filename string, pubkey *rsa.PublicKey) error {
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

func GenerateKeyPair() error {
	privateKey, publicKey, err := GenerateKeyRSA(4096)
	if err != nil {
		fmt.Println("Erreur lors de la génération de la paire de clés:", err)
		return err
	}

	err = security.SavePEMKey("./DUCKY/.ssh/private_key.pem", privateKey)
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde de la clé privée:", err)
		return err
	}

	err = security.SavePEMKeyPublic("./DUCKY/.ssh/public_key.pem", publicKey)
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde de la clé publique:", err)
		return err
	}

	fmt.Println("Paire de clés générée et sauvegardée avec succès.")
	return nil
}

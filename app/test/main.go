package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
)

func encryptECDSA(text string, privateKey *ecdsa.PrivateKey) (string, error) {
	hash := sha256.Sum256([]byte(text))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	signature := append(r.Bytes(), s.Bytes()...)

	encodedSignature := base64.StdEncoding.EncodeToString(signature)

	return encodedSignature, nil
}


func decryptECDSA(text string, publicKey *ecdsa.PublicKey) (bool, error) {

	decodedSignature, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return false, err
	}

	r := new(big.Int)
	s := new(big.Int)
	r.SetBytes(decodedSignature[:len(decodedSignature)/2])
	s.SetBytes(decodedSignature[len(decodedSignature)/2:])

	hash := sha256.Sum256([]byte("Hello, World!"))

	valid := ecdsa.Verify(publicKey, hash[:], r, s)

	return valid, nil
}

func main() {
	privateKey, err := ecdsa.GenerateKey(ecdsa.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Erreur lors de la génération de la clé privée:", err)
		return
	}

	publicKey := &privateKey.PublicKey

	message := "Hello, World!"

	encodedSignature, err := encryptECDSA(message, privateKey)
	if err != nil {
		fmt.Println("Erreur lors du chiffrement:", err)
		return
	}

	fmt.Println("Signature ECDSA:", encodedSignature)

	valid, err := decryptECDSA(encodedSignature, publicKey)
	if err != nil {
		fmt.Println("Erreur lors du déchiffrement:", err)
		return
	}

	if valid {
		fmt.Println("La signature est valide.")
	} else {
		fmt.Println("La signature n'est pas valide.")
	}
}

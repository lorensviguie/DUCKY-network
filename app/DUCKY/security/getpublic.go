package security

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetPublicKey() string {
	publicKeyPath := filepath.Join("./DUCKY/.ssh", "public_key.pem")
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la clé publique :", err)
		panic(0)
	}
	return string(publicKeyBytes)
}

package security

import (
	"fmt"
	"os"
	"path/filepath"
	store "DUCKY/client/structure"
)

func GetPublicKey() string {
	publicKeyPath := filepath.Join(store.KeyPath, "/public.pem")
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la clé publique :", err)
		return "err"
	}
	return string(publicKeyBytes)
}

func GetServeurPublicKey() string {
	publicKeyPath := filepath.Join(store.KeyPath, "/serveurpublickey.pem")
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la clé publique du serveur:", err)
		return "err"
	}
	return string(publicKeyBytes)
}

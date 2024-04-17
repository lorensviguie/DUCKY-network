package bytesreader

import (
	store "DUCKY/client/structure"
	"fmt"
	"os"
	"path/filepath"
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

func GetPrivateKey() string {
	publicKeyPath := filepath.Join(store.KeyPath, "/private.pem")
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la clé privé :", err)
		return "err"
	}
	return string(publicKeyBytes)
}

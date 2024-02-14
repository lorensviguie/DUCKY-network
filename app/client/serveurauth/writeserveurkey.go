package serveurauth

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteToFile(content string) error {
	// Définir le chemin du fichier et le créer s'il n'existe pas
	filePath := filepath.Join(".ssh", "serveurpublickey.pem")

	// Assurer que le répertoire .ssh existe
	err := os.MkdirAll(".ssh", os.ModePerm)
	if err != nil {
		return err
	}

	// Écrire le contenu dans le fichier
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Contenu écrit avec succès dans", filePath)
	return nil
}

package logs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type LogData struct {
	Content string `json:"content"`
}

func LogToFile(fileName string, logContent string) error {
	// Vérifier si le répertoire logs existe, sinon le créer
	logsDir := "./logs/"
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		err = os.Mkdir(logsDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("erreur lors de la création du répertoire logs : %v", err)
		}
	}

	// Chemin complet vers le fichier
	filePath := filepath.Join(logsDir, fileName+".json")

	// Vérifier si le fichier existe, sinon le créer
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("erreur lors de la création du fichier %s : %v", fileName, err)
		}
		defer file.Close()
	}

	// Lire le fichier JSON actuel
	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier %s : %v", fileName, err)
	}
	defer file.Close()

	var logs []LogData
	err = json.NewDecoder(file).Decode(&logs)
	if err != nil && err.Error() != "EOF" {
		return fmt.Errorf("erreur lors de la lecture du fichier JSON %s : %v", fileName, err)
	}

	// Ajouter le nouveau log
	newLog := LogData{Content: string(logContent + time.Now().Format("02/01/2006:15:04"))}
	logs = append(logs, newLog)
	// Réécrire le fichier avec le nouveau log
	file.Seek(0, 0)
	file.Truncate(0)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(logs)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture du fichier JSON %s : %v", fileName, err)
	}

	return nil
}

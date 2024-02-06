package affichage

import (
	"fmt"

	"github.com/fatih/color"
)

func FormatAndDisplay(input string) {
	// Définir les couleurs
	mainColor := color.New(color.FgHiBlue)
	secondaryColor := color.New(color.FgHiMagenta)
	specialColor := color.New(color.FgHiYellow)

	// Appliquer les couleurs à la chaîne
	formattedText := fmt.Sprintf("%s%s%s",
		mainColor.Sprintf("%s", input),
		secondaryColor.Sprintf("\nSecondaire: %s", input),
		specialColor.Sprintf("\nSpécial: %s", input))

	// Afficher le texte formaté
	fmt.Println(formattedText)
}

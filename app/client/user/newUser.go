package user

import (
	security "DUCKY/client/security"
	"fmt"
)

func NewUser() {
	privateKey, publicKey, err := security.GenerateKeyPair(2048)
	if err != nil {
		fmt.Println("Erreur lors de la génération de la paire de clés:", err)
		return
	}

	err = security.SavePEMKey(".ssh/private.pem", privateKey)
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde de la clé privée:", err)
		return
	}

	err = security.SavePEMKeyPublic(".ssh/public.pem", publicKey)
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde de la clé publique:", err)
		return
	}

	fmt.Println("Paire de clés générée et sauvegardée avec succès.")
}

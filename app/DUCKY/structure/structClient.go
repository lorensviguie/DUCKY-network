package structure

import (
	"errors"
	"fmt"
	"net"
)

type ClientInfo struct {
	IP   string
	Conn net.Conn
}

type Authentification struct {
	RandomAuth []byte
	AuthID     string
	Username   string
}

type Session struct {
	Username   string
	SessionID  string
	Conn       net.Conn
	PrivateKey string
}

type Invite struct {
	Username  string
	Room_Name string
}

var Invite_List []Invite
var Sessions []Session

func GetPrivateKeyByUsername(sessions []Session, username string) (string, error) {
	for _, session := range sessions {
		if session.Username == username {
			return session.PrivateKey, nil
		}
	}
	return "", errors.New("Aucune clé privée trouvée pour l'utilisateur " + username)
}

func FindUsernameByConn(conn net.Conn) (string, error) {
	for _, session := range Sessions {
		if session.Conn == conn {
			return session.Username, nil
		}
	}
	return "", fmt.Errorf("Username not found for the given connection")
}

var HelpString = `
Pour parler dans le chat général, utilisez simplement vos messages normalement.
Pour inviter une personne dans une salle, utilisez la commande /invite nom_de_la_salle nom_du_utilisateur.
Pour accepter une invitation, utilisez la commande /accept_invite nom_de_la_salle.
Pour parler dans une salle spécifique, utilisez la commande \nom_de_la_salle suivi de votre message.

Lorsque vous lancez le programme, vous pouvez fournir en argument le chemin vers un fichier contenant vos clés d'utilisateurs. Ceci est utile si vous avez plusieurs utilisateurs.
`

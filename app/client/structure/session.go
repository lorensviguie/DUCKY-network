package structure

type Session struct {
	Username  string
	SessionID []byte
	PublicKey string
}

var ServeurCheck = false
var Authentifier = false
var ServeurAUth []byte
var Sessions Session

var KeyPath = ".ssh"

var HelpString = `
Pour parler dans le chat général, utilisez simplement vos messages normalement.
Pour inviter une personne dans une salle, utilisez la commande /invite nom_de_la_salle nom_du_utilisateur.
Pour accepter une invitation, utilisez la commande /accept_invite nom_de_la_salle.
Pour parler dans une salle spécifique, utilisez la commande \nom_de_la_salle suivi de votre message.

Lorsque vous lancez le programme, vous pouvez fournir en argument le chemin vers un fichier contenant vos clés d'utilisateurs. Ceci est utile si vous avez plusieurs utilisateurs.
`
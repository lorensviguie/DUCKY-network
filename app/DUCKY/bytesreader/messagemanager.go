package bytesreader

import (
	auth "DUCKY/DUCKY/authentification"
	db "DUCKY/DUCKY/database"
	security "DUCKY/DUCKY/security"
	authStorage "DUCKY/DUCKY/structure"
	"bytes"
	"fmt"
	"net"
	"strings"
)

var StorageAuth []authStorage.Authentification

func sendAuthRequest(lines []string) []byte {

	username := StartAuthentificationUser(lines)
	cryptMSG, token, alphaCheck := auth.StartAuthentificationUser(username)
	if alphaCheck == "no" {
		return []byte("Auth Failed please retry")
	}
	nouvelleAuth := authStorage.Authentification{
		RandomAuth: token,
		AuthID:     alphaCheck,
		Username:   username,
	}
	fmt.Println(token)
	StorageAuth = append(StorageAuth, nouvelleAuth)
	return append([]byte("startauthentification\n"+alphaCheck+"\n"), cryptMSG...)
}

func CheckAuth(lines []string, conn net.Conn) []byte {
	alphaCheck := lines[1]
	randomAuth, username := GetRandomAuthByAuthID(alphaCheck)
	returnchack := []byte(lines[2])
	if bytes.Equal(randomAuth, returnchack) {
		privateKey, publicKey, _ := security.GenerateKeyRSA(2048)
		addSession := authStorage.Session{
			Username:   username,
			SessionID:  alphaCheck,
			Conn:       conn,
			PrivateKey: security.PrivateKeyToString(privateKey),
		}
		authStorage.Sessions = append(authStorage.Sessions, addSession)
		messageCrypt, err := security.EncryptMessageWithPublic(db.SelectKeyFromUser(username), (security.PublicKeyToString(publicKey)))
		if err != nil {
			fmt.Println(err)
		}
		return []byte("01\nYou are authentificate Has : \n" + username + "\nWith This ID :\n" + alphaCheck + "\n" + string(messageCrypt))
	} else {
		return []byte("You are not authentificate")
	}
}

func ProveIdentity(lines []string) []byte {
	messagebyte := []byte(strings.Join(lines[1:], "\n"))
	messagedecrypt, err := security.DecryptMessage(security.GetPrivateKey(), messagebyte)
	if err != nil {
		fmt.Println(err)
	}
	return append([]byte("proveidentity\n"), messagedecrypt...)
}

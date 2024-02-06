package bytesreader

import (
	auth "DUCKY/DUCKY/authentification"
	authStorage "DUCKY/DUCKY/structure"
	"bytes"
	"fmt"
)

var StorageAuth []authStorage.Authentification
var Session []authStorage.Session

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

func CheckAuth(lines []string) []byte {
	alphaCheck := lines[1]
	randomAuth, username := GetRandomAuthByAuthID(alphaCheck)
	returnchack := []byte(lines[2])
	fmt.Println(randomAuth)
	fmt.Println(returnchack)
	if bytes.Equal(randomAuth, returnchack) {
		addSession := authStorage.Session{
			Username:  username,
			SessionID: alphaCheck,
		}
		Session = append(Session, addSession)
		return []byte("You are authentificate Has : " + username + "\nWith This ID :\n" + alphaCheck)
	} else {
		return []byte("You are not authentificate")
	}
}

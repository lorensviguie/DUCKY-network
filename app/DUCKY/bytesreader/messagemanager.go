package bytesreader

import (
	auth "DUCKY/DUCKY/authentification"
	authStorage "DUCKY/DUCKY/structure"
	"bytes"
	"fmt"
	"net"
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
		addSession := authStorage.Session{
			Username:  username,
			SessionID: alphaCheck,
			Conn:      conn,
		}
		authStorage.Sessions = append(authStorage.Sessions, addSession)
		return []byte("01\nYou are authentificate Has : \n" + username + "\nWith This ID :\n" + alphaCheck)
	} else {
		return []byte("You are not authentificate")
	}
}

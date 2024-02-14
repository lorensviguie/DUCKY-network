package structure

type Session struct {
	Username  string
	SessionID []byte
}

var ServeurCheck = false
var ServeurAUth []byte
var Sessions []Session

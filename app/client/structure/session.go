package structure

type Session struct {
	Username  string
	SessionID []byte
}

var Sessions []Session

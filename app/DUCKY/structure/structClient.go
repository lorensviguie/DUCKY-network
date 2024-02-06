package structure

import "net"

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
	Username  string
	SessionID string
}

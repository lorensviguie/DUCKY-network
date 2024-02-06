package bytesreader

import (
	"net"
)

func ReadHeaderSize(conn net.Conn) int {
	headerSizeBuf := make([]byte, 1)
	_, err := conn.Read(headerSizeBuf)
	if err != nil {
		return 0
	}
	return int(headerSizeBuf[0])
}



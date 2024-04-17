package bytesreader

import (
	decrypt "DUCKY/client/decrypt"
	"fmt"
	"strings"
)

func MessageStartAuth(lines []string) []byte {
	alphaCheck := []byte(lines[1])
	messagebyte := []byte(strings.Join(lines[2:], "\n"))
	messagedecrypt, err := decrypt.DecryptMessageWithPrivate(GetPrivateKey(), messagebyte)
	if err != nil {
		fmt.Println(err)
	}
	message := []byte(fmt.Sprintf("startauthentification\n%s\n%s", alphaCheck, messagedecrypt))
	return message
}

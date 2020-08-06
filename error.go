package tcp.middleware

import (
	"net"
	"strings"
	"syscall"
)

func IsErrConnReset(err error) bool {
	if ne, ok := err.(*net.OpError); ok {
		return strings.Contains(ne.Err.Error(), syscall.ECONNRESET.Error())
	}
	return false
}

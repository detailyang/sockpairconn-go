package sockpairconn

import (
	"net"
	"os"
)

func newFileConn(fd int) (net.Conn, error) {
	f := os.NewFile(uintptr(fd), "")
	defer f.Close()
	return net.FileConn(f)
}

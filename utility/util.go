package utility

import (
	"net"
	"strconv"
)

func IsPortAvailable(port int) bool {
	address := ":" + strconv.Itoa(port)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return false // Port is busy
	}
	err = ln.Close()
	if err != nil {
		return false
	} // Close the listener to release the port
	return true // Port is available
}

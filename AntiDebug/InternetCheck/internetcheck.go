package InternetCheck

import (
	"net"
	"os"
)

func CheckConnection() {
	_, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		os.Exit(-1)
	}
}

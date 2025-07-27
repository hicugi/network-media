package config

import (
	"os"
	"net"
	"fmt"
)

var SERVER_PORT = (func() string {
	if os.Getenv("PORT") != "" {
		return os.Getenv("PORT")
	}

	return "8000"
})();
var SERVER_IP string = ""

func GetIP() (string, error) {
	if SERVER_IP != "" {
		return SERVER_IP, nil
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				SERVER_IP = ipNet.IP.String()
				return SERVER_IP, nil
			}
		}
	}

	return "", fmt.Errorf("Local IP not found")
}


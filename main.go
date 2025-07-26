package main

import (
	"fmt"
	"net"
	"net/http"
)

const SERVER_PORT = 8000

func getMyIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("Local IP not found")
}

func main() {
	myIp, err := getMyIp()
	if err != nil {
		fmt.Println("Error getting interface addresses:", err)
	}

	fmt.Printf("Starting server on http://%v:%v", myIp, SERVER_PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", SERVER_PORT), nil)
}

package main

import (
	"fmt"
	"net"
	"net/http"

	"medianetwork/internal/Media"
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

type Handler func(w http.ResponseWriter, req *http.Request)
func handleReq(path string, callback Handler) {
	http.HandleFunc(path, (func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("GET: %s\n", path)
		callback(w, req)
	}))
}

func main() {
	myIp, err := getMyIp()
	if err != nil {
		fmt.Println("Error getting interface addresses:", err)
	}

	handleReq("/api/list", media.GetMediaList)

	fmt.Printf("Starting server on http://%v:%v\n", myIp, SERVER_PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", SERVER_PORT), nil)
}

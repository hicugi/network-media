package main

import (
	"fmt"
	"net"
	"net/http"
	"io"
	"os"
	"strings"
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

func getMediaList(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	files, err := os.ReadDir(".")
	if err != nil {
		http.Error(w, "{\"message\":\"Couldn't read dir\"}", http.StatusInternalServerError)
		return
	}

	value := ""
	count := 0

	for _, file := range files {
		name := file.Name()
		format := strings.ToLower(string(name[len(name)-4:]))

		if format != ".mp4" {
			continue
		}

		count++

		if value == "" {
			value = fmt.Sprintf("\"%s\"", name)
			continue
		}

		value = fmt.Sprintf("%s,\"%s\"", value, name)
	}

	res := fmt.Sprintf("{\"items\":[%s]}\n", value);
	fmt.Printf("items found %v res %v\n", count, value);

	io.WriteString(w, res)
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

	handleReq("/api/list", getMediaList)

	fmt.Printf("Starting server on http://%v:%v\n", myIp, SERVER_PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", SERVER_PORT), nil)
}

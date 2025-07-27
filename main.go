package main

import (
	"fmt"
	"net/http"

	"medianetwork/internal/Config"
	"medianetwork/internal/Media"
	"medianetwork/internal/Front"
)

type Handler func(w http.ResponseWriter, req *http.Request)
func handleReq(path string, callback Handler) {
	http.HandleFunc(path, (func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("GET: %s\n", path)
		callback(w, req)
	}))
}

func main() {
	myIp, err := config.GetIP()
	if err != nil {
		fmt.Println("Error getting interface addresses:", err)
	}

	handleReq("/", front.MainPage)
	handleReq("/api/list", media.GetMediaList)
	handleReq("/api/video/play", media.PlayVideo)

	fmt.Printf("Starting server on http://%v:%v\n", myIp, config.SERVER_PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.SERVER_PORT), nil)
}

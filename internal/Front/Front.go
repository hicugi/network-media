package front

import (
	_ "embed"
	"net/http"
	"fmt"
	"os"
	"io"
	"strings"
	"medianetwork/internal/Config"
)

//go:embed index.html
var html string

func MainPage(w http.ResponseWriter, req *http.Request) {
	if os.Getenv("DEV") != "" {
		data, err := os.ReadFile("internal/Front/index.html")
		if err != nil {
			fmt.Printf("Couldn't read the file: %v\n", err)
		}

		html = string(data);
	}

	myIp, err := config.GetIP()
	if err == nil {
		host := fmt.Sprintf("%s:%s", myIp, config.SERVER_PORT)
		html = strings.Replace(html, "{{SERVER_HOST}}", host, 1)
	}

	io.WriteString(w, html)
}

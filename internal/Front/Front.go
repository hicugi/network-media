package front

import (
	_ "embed"
	"net/http"
	"fmt"
	"os"
	"io"
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

	io.WriteString(w, html)
}

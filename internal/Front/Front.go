package front

import (
	"net/http"
	"os"
	"fmt"
	"io"
)

func MainPage(w http.ResponseWriter, req *http.Request) {
	filepath := "internal/Front/index.html"

	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Couldn't read the file: %v\n", err)
	}

	html := string(data);
	io.WriteString(w, html)
}

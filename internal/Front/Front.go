package front

import (
	_ "embed"
	"net/http"
	"io"
)

//go:embed index.html
var html string

func MainPage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, html)
}

package media

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
	"crypto/sha256"
)

func GetMediaList(w http.ResponseWriter, req *http.Request) {
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

func PlayVideo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "video/mp4")

	queryParams := req.URL.Query()
	file := queryParams.Get("s")

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error while reading video file:", err)
		io.WriteString(w, "\r\n")
		return
	}

	w.Header().Del("Content-Length")
	w.Header().Set("Trailer", "X-Content-SHA256, X-Content-Length")


	size, err := io.WriteString(w, string(data))
	if err != nil {
		fmt.Println("Error while writing response body:", err)
		io.WriteString(w, "\r\n")
		return
	}

	sha256 := fmt.Sprintf("%x", sha256.Sum256(data))
	w.Header().Set("X-Content-SHA256", sha256)
	w.Header().Set("X-Content-Length", fmt.Sprintf("%d", size))

	io.WriteString(w, "\r\n")
}

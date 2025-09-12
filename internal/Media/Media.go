package media

import (
	"fmt"
	"net/http"
	"slices"
	"io"
	"os"
	"strings"
	"crypto/sha256"
)

func GetMediaList(w http.ResponseWriter, req *http.Request) {
	FORMATS := []string{"ftypmp42", "ftypqt  ", "ftypisom", "ftypavc1"}

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

		idx := strings.LastIndex(name, ".")
		if idx == -1 {
			continue
		}

		f, err := os.Open(name)
		if err != nil {
			continue
		}

		gap := 4
		buf := make([]byte, 8 + gap)

		n, err := f.Read(buf)
		if err != nil {
			continue
		}

		format := string(buf[gap:n])
		f.Close();

		if !slices.Contains(FORMATS, format) {
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

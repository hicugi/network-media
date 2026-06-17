package media

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type signature struct {
	offset int
	data   []byte
}

var signatures = []signature{
	{4, []byte("ftypmp42")},
	{4, []byte("ftypqt  ")},
	{4, []byte("ftypisom")},
	{4, []byte("ftypavc1")},
	{0, []byte{0x1A, 0x45, 0xDF, 0xA3}}, // MKV
}

func ReadBytes(fileName string, offset, length int) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Couldn't open the file")
	}
	defer f.Close()

	buf := make([]byte, offset+length)
	_, err = f.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read the file")
	}

	return buf[offset : offset+length], nil
}

func matchesSignature(name string) bool {
	for _, sig := range signatures {
		data, err := ReadBytes(name, sig.offset, len(sig.data))
		if err != nil {
			continue
		}
		if bytes.Equal(data, sig.data) {
			return true
		}
	}
	return false
}

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

		idx := strings.LastIndex(name, ".")
		if idx == -1 {
			continue
		}

		if !matchesSignature(name) {
			continue
		}

		info, _ := ReadBytes(name, 0, 255)
		fmt.Printf("%s | %s\n", name, info)

		count++

		if value == "" {
			value = fmt.Sprintf("\"%s\"", name)
			continue
		}

		value = fmt.Sprintf("%s,\"%s\"", value, name)
	}

	res := fmt.Sprintf("{\"items\":[%s]}\n", value)
	fmt.Printf("items found %v res %v\n", count, value)

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

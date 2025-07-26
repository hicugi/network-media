package media

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
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


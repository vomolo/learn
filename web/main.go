package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"as/asciiArt"
)

func main() {
	// File server serves static files over http.   http.Dir  points to the directory from which FileServer will serve the static file
	// http.Handle maps the root url
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// http.Handler is an intefrface that supports receiving and sending back requests over http, and  POST Is one of the requests

	// ascii-art is the end poiint at  which backend handles the request sent by frontend.

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}
		//
		var request struct {
			Banner string `json:"banner"`
			Input  string `json:"input"`
		}
		// decoding the JavascriptObject Notaion Data
		// r.Body is the request body of the incoming http reqrequest
		// jso.NewDecoder  is a decoder taking only r.body as input, i.e taking only JSON data from  the request body

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		// choosing the bannerfile to be mapped// the banner file selected by the user
		fileName := asciiArt.BannerFile(request.Banner)

		// Load the banner map from the file
		bannerMap, err := asciiArt.LoadBannerMap(fileName)
		if err != nil {
			http.Error(w, "Failed to load banner file", http.StatusInternalServerError)
			return
		}

		response := generateASCIIArt(request.Input, bannerMap)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(response))
	})

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// logic for printing the asciii banner files in the required format
func generateASCIIArt(input string, bannerMap map[int][]string) string {
	var str strings.Builder
	lines := make([]string, 8)
	input = strings.ReplaceAll(input, "\r", "\n")
	arr := strings.Split(input, "\n")
	for _, v := range arr {
		for _, char := range v {
			banner, exists := bannerMap[int(char)]
			if !exists {
				banner = bannerMap[32] // Fallback to space if character not found
			}
			for i := 0; i < 8; i++ {
				lines[i] += banner[i]
			}
		}
		str.WriteString(strings.Join(lines, "\n"))
		str.WriteString("\n")
		lines = make([]string, 8)
	}
	return str.String()
}

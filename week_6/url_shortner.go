package week_6

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

var urls = make(map[string]string)

func generateKey() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	key := generateKey()
	urls[key] = longURL

	shortURL := fmt.Sprintf("http://127.0.0.1:8080/%s", key)
	fmt.Fprintf(w, "Shortened URL: %s", shortURL)
}

func AllShortenedURLsHandler(w http.ResponseWriter, r *http.Request) {
	for key, longURL := range urls {
		fmt.Fprintf(w, "Key: %s, Long URL: %s -> http://localhost:8080/%s\n", key, longURL,key)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/")
	longURL, ok := urls[key]
	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

func StartURLShortener() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/all", AllShortenedURLsHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("URL Shortener server started on :8080")
	http.ListenAndServe(":8080", nil)
}
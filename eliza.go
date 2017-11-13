package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func elizaSessionHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "elizaSession.html")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/elizaSession", elizaSessionHandler)
	http.ListenAndServe(":8080", nil)
}

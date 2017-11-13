package main

import (
	"net/http"
)

type message struct {
	Header    string
	UserGuess int
	Message   string
	Won       bool
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

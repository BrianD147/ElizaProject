package main

import (
	"html/template"
	"net/http"
)

type UsernameData struct {
	Username string
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, UsernameData{Username: "Test"})
}

func elizaSessionHandler(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//name := r.Form["userNameInput"]
	//fmt.Println(name)
	t, _ := template.ParseFiles("elizaSession.html")
	t.Execute(w, UsernameData{Username: "Test"})
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/elizaSession", elizaSessionHandler)
	http.ListenAndServe(":8080", nil)
}

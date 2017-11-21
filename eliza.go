package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type UsernameData struct {
	Username string
}

func handler(w http.ResponseWriter, r *http.Request) {

	message := r.URL.Query().Get("messageSubmit")
	//r.ParseForm()
	//name := r.Form["userNameInput"]
	fmt.Println(message)
	t, _ := template.ParseFiles("elizaSession.html")
	t.Execute(w, UsernameData{Username: "Test"})
}

func Eliza(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Eliza response")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//fmt.Println("People say I look like both my mother and father.")
	//fmt.Println(elizaResponse("People say I look like both my mother and father."))
	fmt.Println()

	http.HandleFunc("/", handler)
	http.HandleFunc("/Eliza", Eliza)
	http.ListenAndServe(":8080", nil)
}

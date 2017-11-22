package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("elizaSession.html")
	t.Execute(w, nil)
}

func Eliza(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("value")

	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, message); matched {
		response := []string{"Tell me more about your father", "I am your father!", "I have a father too, he's very handsome"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
	} else if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, message); matched {
		response := []string{"Tell me more about your mother", "I don't have a mother!"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
	} else if matched, _ := regexp.MatchString(`(?i).*\bsister\b.*`, message); matched {
		response := []string{"Tell me more about your sister", "I have a sister, her name is Siri!"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
	} else if matched, _ := regexp.MatchString(`(?i).*\bbrother\b.*`, message); matched {
		response := []string{"Tell me more about your brother", "I don't have a brother!"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
	} else {
		response := []string{"I’m not sure what you’re trying to say. Could you explain it to me?", "How does that make you feel?", "Why do you say that?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", handler)
	http.HandleFunc("/Eliza", Eliza)
	http.ListenAndServe(":8080", nil)
}

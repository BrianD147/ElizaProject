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
	} else if matched, _ := regexp.MatchString(`(?i).*i remember ?(.*)`, message); matched {
		response := []string{"What else do you recollect?", "That's a good memory!"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
	} else if matched, _ := regexp.MatchString(`(?i).*\bsorry\b.*`, message); matched {
		response := []string{"Please don't apologise.", "Apologies are not necessary.", "I've told you that apologies are not required."}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
	} else if matched, _ := regexp.MatchString(`(?i).*\bdream\b.*`, message); matched {
		response := []string{"What does that dream suggest to you?", "Do you dream often?", "What persons appear in your dreams?", "Do you believe that dreams have something to do with your problems?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
	} else if matched, _ := regexp.MatchString(`(?i).*\bperhaps\b.*`, message); matched {
		response := []string{"You don't seem quite certain.", "Why the uncertain tone?", "Can't you be more positive?", "You aren't sure?", "Don't you know?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
	} else {
		response := []string{"I’m not sure what you’re trying to say. Could you explain it to me?", "I don't quite get you!", "What do you mean?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", handler)
	http.HandleFunc("/Eliza", Eliza)
	http.ListenAndServe(":8080", nil)
}

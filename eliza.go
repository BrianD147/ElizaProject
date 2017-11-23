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
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, message); matched {
		response := []string{"Tell me more about your mother", "I don't have a mother!"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bsister\b.*`, message); matched {
		response := []string{"Tell me more about your sister", "I have a sister, her name is Siri!"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bbrother\b.*`, message); matched {
		response := []string{"Tell me more about your brother", "I don't have a brother!"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*i remember ?(.*)`, message); matched {
		response := []string{"What else do you recollect?", "That's a good memory!"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bsorry\b.*`, message); matched {
		response := []string{"Please don't apologise.", "Apologies are not necessary.", "I've told you that apologies are not required."}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bdream\b.*`, message); matched {
		response := []string{"What does that dream suggest to you?", "Do you dream often?", "What persons appear in your dreams?", "Do you believe that dreams have something to do with your problems?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bperhaps\b.*`, message); matched {
		response := []string{"You don't seem quite certain.", "Why the uncertain tone?", "Can't you be more positive?", "You aren't sure?", "Don't you know?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bno\b.*`, message); matched {
		response := []string{"You'll have to elaborate.", "Definitely no?", "Can you tell me why not?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\byes\b.*`, message); matched {
		response := []string{"You seem quite sure.", "OK, but can you elaborate a bit?", "Can you tell me why?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bbecause\b.*`, message); matched {
		response := []string{"Is that the real reason?", "What other reasons come to mind?", "Does that reason apply to anything else?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bsorry\b.*`, message); matched {
		response := []string{"There are many times when no apology is needed.", "What feelings do you have when you apologize?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bhello\b.*`, message); matched {
		response := []string{"Hello, I'm glad we have a chance to chat.", "Hi there, how are you today?", "Hello, how are you feeling today?", "Alright, mate?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	} else if matched, _ := regexp.MatchString(`(?i).*\bfriend\b.*`, message); matched {
		response := []string{"Tell me more about your friends.", "When you think of a friend, what comes to mind?", "Did you have many childhood friends?"}
		fmt.Fprintf(w, response[rand.Intn(len(response))])
		return
	}

	re := regexp.MustCompile(`(?i)^\s*I need ([^.!?]*)[.!?]*\s*$`)
	if re.MatchString(message) {
		chosenResponse := []string{"Why do you need $1?", "Would it really help you to get $1?", "Are you sure you need $1?"}
		response := re.ReplaceAllString(message, chosenResponse[rand.Intn(len(chosenResponse))])
		fmt.Fprintf(w, response)
	}

	re = regexp.MustCompile(`(?i)^\s*I can'?t ([^.!?]*)[.!?\s]*$`)
	if re.MatchString(message) {
		chosenResponse := []string{"How do you know you can't $1?", "Perhaps you could $1 if you tried", "What would it take for you to $1?"}
		response := re.ReplaceAllString(message, chosenResponse[rand.Intn(len(chosenResponse))])
		fmt.Fprintf(w, response)
	}

	re = regexp.MustCompile(`(?i)^\s*(?:I am|I'm) ([^.!?]*)[.!?\s]*$`)
	if re.MatchString(message) {
		chosenResponse := []string{"Did you come to me because you are $1?", "How long have you been $1?", "How do you feel about being $1?", "How does being $1 make you feel?", "Do you enjoy being $1?", "Why do you say that you're $1?", "Why do you think you're $1?"}
		response := re.ReplaceAllString(message, chosenResponse[rand.Intn(len(chosenResponse))])
		fmt.Fprintf(w, response)
	}

	re = regexp.MustCompile(`(?i)^\s*How .*$`)
	if re.MatchString(message) {
		chosenResponse := []string{"How do you suppose?", "Perhaps you can answer your own question.", "What is it you're really asking?"}
		response := re.ReplaceAllString(message, chosenResponse[rand.Intn(len(chosenResponse))])
		fmt.Fprintf(w, response)
	}

	re = regexp.MustCompile(`(?i)^\s*Because ([^.!?]*)[.!?\s]*$`)
	if re.MatchString(message) {
		chosenResponse := []string{"Is that the real reason?", "What other reasons come to mind?", "Does that reason apply to anything else?", "If $1, what else must be true?"}
		response := re.ReplaceAllString(message, chosenResponse[rand.Intn(len(chosenResponse))])
		fmt.Fprintf(w, response)
	}

	re = regexp.MustCompile(`(?i)^\s*It is ([^.!?]*)[.!?\s]*$`)
	if re.MatchString(message) {
		chosenResponse := []string{"You seem very certain.", "If I told you that it probably isn't $1, how would you feel?"}
		response := re.ReplaceAllString(message, chosenResponse[rand.Intn(len(chosenResponse))])
		fmt.Fprintf(w, response)
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

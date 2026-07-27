package api

import (
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleHello)
	http.HandleFunc("/api/questions/", handleQuestions)
	http.HandleFunc("/api/answers/", handleAnswers)
	http.HandleFunc("/api/votes/", handleVotes)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to App Engine")
}

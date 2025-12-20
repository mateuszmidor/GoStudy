package main

import (
	"log"
	"net/http"
)

func main() {
	// setup endpoints
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/api/auth/login", login)
	http.HandleFunc("/api/auth/logout", logout)

	// run server
	log.Println("listening at :8080")
	http.ListenAndServe(":8080", nil)
}

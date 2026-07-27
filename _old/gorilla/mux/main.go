package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const address = "localhost:8000"

func main() {
	r := mux.NewRouter()

	// only handle index page
	r.HandleFunc("/", indexHandler)

	// handle all subpaths under "/files"
	r.PathPrefix("/files").HandlerFunc(catchAllHandler)

	// handle subpaths under "/auth", but use also authorization middleware
	sub := r.PathPrefix("/auth").Subrouter()
	sub.Use(printBasicAuth)
	sub.PathPrefix("/").HandlerFunc(catchAllHandler)

	// Bind to a port and pass our router in
	log.Printf("listening at %s", address)
	log.Fatal(http.ListenAndServe(address, r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index!\n"))
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
}

func printBasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		log.Printf("%t - %s:%s", ok, user, pass)
		next.ServeHTTP(w, r)
	})
}

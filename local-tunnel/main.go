package main

import (
	"log"
	"net/http"
)

const Address = ":8080"

func main() {
	http.HandleFunc("/", handler)
	log.Println("serving at", Address)
	log.Fatal(http.ListenAndServe(Address, nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)
	response := "Hello world!"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

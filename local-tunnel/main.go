package main

import (
	"fmt"
	"log"
	"net/http"
)

const Address = "0.0.0.0:33000"

var counter int

func main() {
	http.HandleFunc("/", handler)
	log.Println("serving at", Address)
	log.Fatal(http.ListenAndServe(Address, nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	log.Printf("%d - a request came!", counter)
	response := fmt.Sprintf("%d - Hello world!", counter)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

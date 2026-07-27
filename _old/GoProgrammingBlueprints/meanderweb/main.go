package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8081", "Adres aplikacji internetowej")
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	log.Println("Aplikacja internetowa jest dostępna pod adresem:", *addr)
	http.ListenAndServe(*addr, mux)
}

package api

import (
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleHello)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to App Engine")
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "create.html")
	})

	fmt.Println("Starting server at port 8888")
	if err := http.ListenAndServeTLS(":8888", "./localhost/cert.pem", "./localhost/key.pem", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

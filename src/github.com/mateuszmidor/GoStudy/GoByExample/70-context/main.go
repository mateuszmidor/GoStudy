package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

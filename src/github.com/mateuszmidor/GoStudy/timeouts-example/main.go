package main

import (
	"fmt"
	"net/http"
	"time"
)

func longrequestHandler(w http.ResponseWriter, r *http.Request) {
	//wait 3 sec and handle request OR handle request cancellation; whatever happens first
	select {
	case <-time.NewTimer(3 * time.Second).C:
		println("longrequestHandler: request handled")
		w.WriteHeader(200)
	case <-r.Context().Done():
		println("longrequestHandler: request cancelled by client side")
		w.WriteHeader(499) // nginx custom code but fits here: client closed connection
	}
}

func createHTTPClientWithTimeout() *http.Client {
	return &http.Client{Timeout: time.Second}
}

func createHTTPServerWithTimeouts() *http.Server {
	return &http.Server{
		Addr:         ":8080",
		Handler:      nil, // use DefaultServeMux
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	}
}

func main() {
	http.HandleFunc("/longrequest", longrequestHandler)
	server := createHTTPServerWithTimeouts()
	go server.ListenAndServe()

	// set client timeout to 1 sec while request processing is at least 3 sec
	// this should generate timeout on client side and request cancellation on server side
	client := createHTTPClientWithTimeout()
	_, err := client.Get("http://localhost:8080/longrequest")
	fmt.Printf("Error: %v\n", err)
}

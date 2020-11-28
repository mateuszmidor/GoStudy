package main

import (
	"fmt"
	"net/http"
	"path"
)

// CA = Certificate Authority; entity that issues certificates
const caDir = "../../../cert/minica"

var serverKeyFile = path.Join(caDir, "localhost/key.pem")
var serverCertFile = path.Join(caDir, "localhost/cert.pem")

func allHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from TLS HTTP server!")
	for k, v := range r.Header {
		fmt.Fprintf(w, "%+30s  %s\n", k, v)
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("TLS HTTP Server listening on port 9000")
	http.HandleFunc("/", allHandler)
	panicOnError(http.ListenAndServeTLS(":9000", serverCertFile, serverKeyFile, nil))
}

package main

import (
	"fmt"
	"net/http"
	"path"
)

const serverCertsDir = "../../../cert/minica/localhost"

var keyFile = path.Join(serverCertsDir, "key.pem")
var certFile = path.Join(serverCertsDir, "cert.pem")

func allHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from TLS server!")
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
	fmt.Println("HTTPS Server listening on port 9000")
	http.HandleFunc("/", allHandler)
	panicOnError(http.ListenAndServeTLS(":9000", certFile, keyFile, nil))
}

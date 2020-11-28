package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

// CA = Certificate Authority; entity that issues certificates
const caDir = "../../../cert/minica"

// CA Cert is needed to trust certificates issued by this specific CA. In this example, client will use such certificate
var caCertFile = path.Join(caDir, "minica.pem")
var serverKeyFile = path.Join(caDir, "localhost/key.pem")
var serverCertFile = path.Join(caDir, "localhost/cert.pem")

func getTLSConfig() *tls.Config {
	// tell the server: trust certificates issued by this Certificate Authority (here: minica)
	pemData, err := ioutil.ReadFile(caCertFile)
	panicOnError(err)
	caCert := x509.NewCertPool()
	caCert.AppendCertsFromPEM(pemData)

	tlsConfig := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert, // require certificate from client
		ClientCAs:  caCert,                         // trust client certificates issued by this CA
		MinVersion: tls.VersionTLS13,
		MaxVersion: tls.VersionTLS13,
	}

	return tlsConfig
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Mutual TLS (mTLS) HTTP server!")
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
	fmt.Println("mTLS HTTP Server listening on port 9000")

	// create mTLS HTTP server
	tlsConfig := getTLSConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/", allHandler)
	emptyNextProto := map[string]func(*http.Server, *tls.Conn, http.Handler){}
	server := http.Server{
		Addr:         ":9000",
		Handler:      mux,
		TLSConfig:    tlsConfig,
		TLSNextProto: emptyNextProto,
	}

	// serve mTLS HTTP clients
	panicOnError(server.ListenAndServeTLS(serverCertFile, serverKeyFile))
}

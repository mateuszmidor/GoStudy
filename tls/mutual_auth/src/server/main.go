package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const caCertPool = "../../minica/cacert.crt"

func parseCert(certFile, keyFile string) (cert tls.Certificate, err error) {
	cert, err = tls.LoadX509KeyPair(certFile, keyFile)
	return
}

// configure and create a tls.Config instance using the provided cert, key, and ca cert files.
func configureTLS(certFile, keyFile, caCertFile string) (tlsConfig *tls.Config, err error) {

	c, err := parseCert(certFile, keyFile)
	if err != nil {
		return
	}

	ciphers := []uint16{
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	}

	certPool := x509.NewCertPool()
	buf, err := ioutil.ReadFile(caCertFile)
	panicOnError(err)

	if !certPool.AppendCertsFromPEM(buf) {
		log.Fatalln("Failed to parse truststore")
	}

	tlsConfig = &tls.Config{
		CipherSuites:             ciphers,
		ClientAuth:               tls.RequireAndVerifyClientCert,
		PreferServerCipherSuites: true,
		RootCAs:                  certPool,
		ClientCAs:                certPool,
		Certificates:             []tls.Certificate{c},
	}

	return
}

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
	panicOnError(http.ListenAndServeTLS(":9000", "../../cert/minica/localhost/cert.pem", "../../cert/minica/localhost/key.pem", nil))
}

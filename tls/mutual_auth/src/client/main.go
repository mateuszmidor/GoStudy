package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

// CA = Certificate Authority; entity that issues certificates
const caDir = "../../../cert/minica"

// CA Cert is needed to trust certificates issued by this specific CA. In this example, server will use such certificate
var caCertFile = path.Join(caDir, "minica.pem")
var clientKeyFile = path.Join(caDir, "clientcert/key.pem")
var clienCertFile = path.Join(caDir, "clientcert/cert.pem")

func getTLSConfig() *tls.Config {
	// tell the server: this is the list of ciphering suites I understand
	supportedCipherSuites := []uint16{
		tls.TLS_RSA_WITH_RC4_128_SHA,
		tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	}

	// tell the client: trust certificates issued by this Certificate Authority (here: minica)
	pemData, err := ioutil.ReadFile(caCertFile)
	panicOnError(err)
	caCert := x509.NewCertPool()
	caCert.AppendCertsFromPEM(pemData)

	// load client certificate
	clientCert, err := tls.LoadX509KeyPair(clienCertFile, clientKeyFile)
	panicOnError(err)

	tlsConfig := &tls.Config{
		CipherSuites:             supportedCipherSuites,
		PreferServerCipherSuites: true,
		RootCAs:                  caCert,                        // trust server certificates issued by this CA
		Certificates:             []tls.Certificate{clientCert}, // provide this certificate to server
		MinVersion:               tls.VersionTLS13,
		MaxVersion:               tls.VersionTLS13,
	}

	return tlsConfig
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// create HTTPS client
	tlsConfig := getTLSConfig()
	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{Transport: tr}

	// call HTTPS server
	resp, err := client.Get("https://localhost:9000")
	panicOnError(err)
	defer resp.Body.Close()

	// print server response to stdout
	_, err = io.Copy(os.Stdout, resp.Body)
	panicOnError(err)

	fmt.Println()
}

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	mTLSConfig := &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_RC4_128_SHA,
			tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		},
	}

	mTLSConfig.PreferServerCipherSuites = true
	mTLSConfig.MinVersion = tls.VersionTLS13
	mTLSConfig.MaxVersion = tls.VersionTLS13

	certs := x509.NewCertPool()

	pemData, err := ioutil.ReadFile("../../cert/minica/minica.pem")
	panicOnError(err)

	certs.AppendCertsFromPEM(pemData)
	mTLSConfig.RootCAs = certs

	tr := &http.Transport{
		TLSClientConfig: mTLSConfig,
	}

	c := &http.Client{Transport: tr}

	resp, err := c.Get("https://localhost:9000")
	panicOnError(err)
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	panicOnError(err)

	fmt.Println()
}

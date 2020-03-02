// Project: http resource fetcher, like wget/curl
// Usage: go run . http://golang.org
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	improvedFetch()
	// for _, url := range os.Args[1:] {
	// 	resp, err := http.Get(url)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	b, err := ioutil.ReadAll(resp.Body)
	// 	resp.Body.Close()
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch: error reading %s: %v\n", url, err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Print(string(b))
	// }
}

func improvedFetch() {
	for _, url := range os.Args[1:] {
		prefixedURL := addHTTPPrefixIfMissing(url)
		resp := httpGetOrExit(prefixedURL)
		ioCopyOrExit(os.Stdout, resp.Body)
		fmt.Println("HTTP STATUS: ", resp.Status)
	}
}

func httpGetOrExit(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	return resp
}

func ioCopyOrExit(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: error copying to stdout %v\n", err)
		os.Exit(1)
	}
}

func addHTTPPrefixIfMissing(url string) string {
	const HTTP = "http://"
	if strings.HasPrefix(url, HTTP) {
		return url
	}
	return HTTP + url
}

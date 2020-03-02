// Project: parallel http resource fetcher, like wget/curl but many resoures in parallel using go routines
// Usage: go run . http://golang.org http://wp.pl http://tvn24.pl
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	// fetch pages
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	// print results
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("Total of %0.2f seconds passed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f seconds\t%7d bytes\t%s", secs, nbytes, url)
}

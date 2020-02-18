// Project: go routined web crawler
// Usage: go run . http://www.golang.org
package main

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"
)

func main() {
	worklist := make(chan []string)
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// do something with item and return strings, eg url -> http links
type processItem func(item string) []string

// worklist contains urls, processItem returns sub-urls for given url
func breadthFirst(f processItem, worklist []string) {

}

// tokens is used to reduce num of parallel goroutines so we dont kill the system :)
var tokens = make(chan struct{}, 20)

var urls uint64

func crawl(url string) []string {
	val := atomic.AddUint64(&urls, 1)
	fmt.Printf("%d. %s\n", val, url)
	tokens <- struct{}{}
	list, err := ExtractLinksFromPage(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}
	return list
}

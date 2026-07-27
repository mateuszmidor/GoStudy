// Project: go routined web crawler with limited depth
// Usage: go run . http://www.golang.org
package main

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"
)

// Links holds the http links and search depth
type Links struct {
	links []string
	depth int
}

func main() {
	const maxDepth = 5 // 0 means only starting page

	worklist := make(chan Links)
	go func() {
		worklist <- Links{os.Args[1:], 0}
	}()

	seen := make(map[string]bool)
	for linksDepth := range worklist {
		if linksDepth.depth > maxDepth {
			continue
		}
		for _, link := range linksDepth.links {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- Links{crawl(link), linksDepth.depth + 1}
				}(link)
			}
		}

		// if len(worklist) == 0 {
		// 	close(worklist)
		// }
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

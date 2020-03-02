// Project: crawl links from provided links
// Usage: wget www.wp.pl -O - -q | go run .
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

// do something with item and return strings, eg url -> http links
type processItem func(item string) []string

// worklist contains urls, processItem returns sub-urls for given url
func breadthFirst(f processItem, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := ExtractLinksFromPage(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

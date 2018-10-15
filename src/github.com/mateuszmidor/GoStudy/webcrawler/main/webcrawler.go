// Based on https://tour.golang.org/concurrency/10

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

// thread safe list of visited urls to avoid repetitions
type VisitedUrls struct {
	urls map[string]bool
	mux  sync.Mutex
}

// if no such url visited yet, add it to the list and return true, else return false
func (visited *VisitedUrls) tryVisitUrl(url string) bool {
	visited.mux.Lock()
	defer visited.mux.Unlock()

	if visited.urls == nil {
		visited.urls = make(map[string]bool)
	}

	if _, ok := visited.urls[url]; ok {
		return false
	}

	visited.urls[url] = true
	return true
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, visitedUrls *VisitedUrls) {
	if depth <= 0 {
		return
	}

	if visitedUrls.tryVisitUrl(url) == false {
		return
	}

	html, urls, err := fetcher.Fetch(url)
	if err != nil {
		// fmt.Println(err)
		return
	}

	fmt.Printf("found: %s - %s\n", url, extractTitle(html))

	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(item_url string) {
			Crawl(item_url, depth-1, fetcher, visitedUrls)
			wg.Done()
		}(u)
	}
	wg.Wait()
	return
}

type SerialFetcher struct {
}

// fetch actual page at given url
func (f SerialFetcher) Fetch(url string) (string, []string, error) {
	html, err := fetchPage(url)
	return html, extractUrls(html), err
}

// fetch html page
func fetchPage(url string) (html string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	return string(bytes), err
}

// extract urls from html
func extractUrls(html string) []string {
	REGEXP := `<a href="(http.+?)"`
	re := regexp.MustCompile(REGEXP)
	matches := re.FindAllStringSubmatch(html, -1)
	var results []string
	for _, submatches := range matches {
		if len(submatches) > 1 {
			results = append(results, submatches[1])
		}
	}
	return results
}

// extract page title from html
func extractTitle(html string) string {
	REGEXP := `<title>(.+)</title>`
	re := regexp.MustCompile(REGEXP)
	matches := re.FindStringSubmatch(html)
	if matches != nil && len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func main() {
	var myFetcher SerialFetcher
	var visitedUrls VisitedUrls
	Crawl("http://mateuszmidor.com", 2, myFetcher, &visitedUrls)
}

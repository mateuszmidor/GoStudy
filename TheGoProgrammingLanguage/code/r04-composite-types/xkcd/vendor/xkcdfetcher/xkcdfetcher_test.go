package xkcdfetcher

import (
	"fetcher"
	"fmt"
	"sharedkernel"
	"testing"
)

var comicBook1 = sharedkernel.ComicBook{Title: "Dzień Śmiechały", Transcription: "Kajkoszowy dzień żartów", URL: "", Number: 1}
var comicBook2 = sharedkernel.ComicBook{Title: "Szkoła latania", Transcription: "Mirmił robi PPL", URL: "", Number: 2}
var comicBook3 = sharedkernel.ComicBook{Title: "Wielki Turniej", Transcription: "Zawody na najwaleczniejszego barana", URL: "", Number: 3}

var comicBook1URL = getURLForComicBookNo(1)
var comicBook2URL = getURLForComicBookNo(2)
var comicBook3URL = getURLForComicBookNo(3)
var latestComicBookURL = getURLForLatestComicBook()

var urlToComicBook = map[string]sharedkernel.ComicBook{
	comicBook1URL:      comicBook1,
	comicBook2URL:      comicBook2,
	comicBook3URL:      comicBook3,
	latestComicBookURL: comicBook3,
}

// Provides canned ComicBook for given URL
func cannedSingleComicBookFetcher(url string) (*sharedkernel.ComicBook, error) {
	comicBook, ok := urlToComicBook[url]
	if !ok {
		return nil, fmt.Errorf("No comic book at %q", url)
	}
	return &comicBook, nil
}

func TestStartFetchingComicBooks(t *testing.T) {
	const expectedComicBooksAvailable = 3

	f := NewXkcdFetcher(cannedSingleComicBookFetcher)
	c := make(fetcher.ComicBookChannel)

	numComicBooksAvailable, err := f.StartFetchingComicBooks(c)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if numComicBooksAvailable != expectedComicBooksAvailable {
		t.Fatalf("Actual number of available comic books != expected: %d != %d", numComicBooksAvailable, expectedComicBooksAvailable)
	}

	if fetched := <-c; fetched == nil {
		t.Errorf("Error fetching comic book1")
	} else if *fetched != comicBook1 {
		t.Errorf("Fetched comic book != expected: %v != %v", fetched, comicBook1)
	}

	if fetched := <-c; fetched == nil {
		t.Errorf("Error fetching comic book2")
	} else if *fetched != comicBook2 {
		t.Errorf("Fetched comic book != expected: %v != %v", fetched, comicBook2)
	}

	if fetched := <-c; fetched == nil {
		t.Errorf("Error fetching comic book3")
	} else if *fetched != comicBook3 {
		t.Errorf("Fetched comic book != expected: %v != %v", fetched, comicBook3)
	}
}

package fetcher

import (
	"sharedkernel"
)

// ComicBookChannel is where the fetched comic books are pushed to
type ComicBookChannel chan *sharedkernel.ComicBook

// Fetcher is interface for online comic books fetcher
type Fetcher interface {
	// StartFetchingComicBooks returns num comic books available for fetching
	StartFetchingComicBooks(c ComicBookChannel) (int, error)
}

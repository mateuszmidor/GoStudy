package infrastructure

import (
	"sharedkernel"
)

// BrowserPort allows outer world talk to the Browser
type BrowserPort interface {
	LoadComicBooks(comicBooks sharedkernel.ComicBooks)
	NumAvailableComicBooks() uint
	GotoComicBookNumber(no uint) bool
	GotoNextComicBook() bool
	GotoPrevComicBook() bool
	GotoLatestComicBook() bool
	GetCurrentComicBook() *sharedkernel.ComicBook
}

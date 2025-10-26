package browser

import (
	"browser/domain"
	"sharedkernel"
)

// Root implements infrastructure.BrowserPort
type Root struct {
	browser *domain.Browser
}

// NewRoot is Root constructor
func NewRoot() *Root {
	return &Root{domain.NewBrowser()}
}

// LoadComicBooks implements BrowserPort.LoadComicBooks
func (r *Root) LoadComicBooks(comicBooks sharedkernel.ComicBooks) {
	r.browser.LoadComicBooks(comicBooks)
}

// NumAvailableComicBooks implements BrowserPort.NumAvailableComicBooks
func (r *Root) NumAvailableComicBooks() uint {
	return r.browser.NumAvailableComicBooks()
}

// GotoComicBookNumber implements BrowserPort.GotoComicBookNumber
func (r *Root) GotoComicBookNumber(no uint) bool {
	return r.browser.GotoComicBookNumber(no)
}

// GotoNextComicBook implements GotoNextComicBook.GotoNextComicBook
func (r *Root) GotoNextComicBook() bool {
	return r.browser.GotoNextComicBook()
}

// GotoPrevComicBook implements BrowserPort.GotoPrevComicBook
func (r *Root) GotoPrevComicBook() bool {
	return r.browser.GotoPrevComicBook()
}

// GotoLatestComicBook implements BrowserPort.GotoLatestComicBook
func (r *Root) GotoLatestComicBook() bool {
	return r.browser.GotoLatestComicBook()
}

// GetCurrentComicBook implements BrowserPort.GetCurrentComicBook
func (r *Root) GetCurrentComicBook() *sharedkernel.ComicBook {
	return r.browser.GetCurrentComicBook()
}

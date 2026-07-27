package domain

import "sharedkernel"

// Browser implements the rules
type Browser struct {
	comicBooks            sharedkernel.ComicBooks
	currentComicBookIndex int
}

// NewBrowser is Browser constructor
func NewBrowser() *Browser {
	return &Browser{sharedkernel.ComicBooks{}, -1}
}

// NumAvailableComicBooks implements BrowserPort.NumAvailableComicBooks
func (b *Browser) NumAvailableComicBooks() uint {
	return uint(len(b.comicBooks))
}

// LoadComicBooks implements BrowserPort.LoadComicBooks
func (b *Browser) LoadComicBooks(comicBooks sharedkernel.ComicBooks) {
	b.comicBooks = comicBooks
	b.currentComicBookIndex = -1 // invalidate
}

// GotoComicBookNumber implements BrowserPort.GotoComicBookNumber
func (b *Browser) GotoComicBookNumber(no uint) bool {
	if no >= b.NumAvailableComicBooks() {
		return false
	}
	b.currentComicBookIndex = int(no)
	return true
}

// GotoNextComicBook implements GotoNextComicBook.GotoNextComicBook
func (b *Browser) GotoNextComicBook() bool {
	if uint(b.currentComicBookIndex+1) >= b.NumAvailableComicBooks() {
		return false
	}
	b.currentComicBookIndex++
	return true
}

// GotoPrevComicBook implements BrowserPort.GotoPrevComicBook
func (b *Browser) GotoPrevComicBook() bool {
	if int(b.currentComicBookIndex)-1 < 0 {
		return false
	}
	b.currentComicBookIndex--
	return true
}

// GotoLatestComicBook implements BrowserPort.GotoLatestComicBook
func (b *Browser) GotoLatestComicBook() bool {
	if b.NumAvailableComicBooks() == 0 {
		return false
	}
	b.currentComicBookIndex = int(b.NumAvailableComicBooks()) - 1
	return true
}

// GetCurrentComicBook implements BrowserPort.GetCurrentComicBook
func (b *Browser) GetCurrentComicBook() *sharedkernel.ComicBook {
	if b.currentComicBookIndex == -1 {
		return nil
	}
	return &b.comicBooks[b.currentComicBookIndex]
}

package ui

import (
	"sharedkernel"
)

// Handler is receier of UI commands
type Handler interface {
	OnPrevComicBook()
	OnNextComicBook()
	OnGotoComicBook(number uint)
}

// UI is interface for user interface
type UI interface {
	Run()
	DisplayStatus(status string)
	DisplayComicBook(comicBook *sharedkernel.ComicBook)
}

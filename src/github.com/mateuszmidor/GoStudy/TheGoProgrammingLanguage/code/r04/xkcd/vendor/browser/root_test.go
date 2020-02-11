package browser

import (
	"sharedkernel"
	"testing"
)

var comicBook1 = sharedkernel.ComicBook{Title: "Dzień Śmiechały", Transcription: "Kajkoszowy dzień żartów", URL: "", Number: 1}
var comicBook2 = sharedkernel.ComicBook{Title: "Szkoła latania", Transcription: "Mirmił robi PPL", URL: "", Number: 2}
var comicBook3 = sharedkernel.ComicBook{Title: "Wielki Turniej", Transcription: "Zawody na najwaleczniejszego barana", URL: "", Number: 3}
var testComicBooks = sharedkernel.ComicBooks{comicBook1, comicBook2, comicBook3}

func TestGotoLatestReturnsFalseWhenNoComicBooksLoaded(t *testing.T) {
	b := NewRoot()
	if b.GotoLatestComicBook() == true {
		t.Error("GotoLatestComicBook should return false when no comic books are loaded")
	}
}

func TestGotoLatestReturnsTrueWhenSomeComicBooksLoaded(t *testing.T) {
	b := NewRoot()
	b.LoadComicBooks(testComicBooks)
	if b.GotoLatestComicBook() == false {
		t.Error("GotoLatestComicBook should return true when some comic books are loaded")
	}
}

func TestGotoComicBookNumbertReturnsFalseWhenNoComicBooksLoaded(t *testing.T) {
	b := NewRoot()
	if b.GotoComicBookNumber(0) == true {
		t.Error("GotoComicBookNumber should return false when no comic books are loaded")
	}
}

func TestGotoComicBookNumbertReturnsTrueWhenSomeComicBooksLoaded(t *testing.T) {
	b := NewRoot()
	b.LoadComicBooks(testComicBooks)
	if b.GotoComicBookNumber(0) == false {
		t.Error("GotoComicBookNumber should return true when some comic books are loaded")
	}
}

func TestGotoNextComicBookReturnsFalseWhenNoComicBooksLoaded(t *testing.T) {
	b := NewRoot()
	if b.GotoNextComicBook() == true {
		t.Error("GotoNextComicBook should return false when no next comic book available")
	}
}

func TestGotoNextComicBookReturnsTrueWhenSomeComicBooksLoaded(t *testing.T) {
	b := NewRoot()
	b.LoadComicBooks(testComicBooks)
	if b.GotoNextComicBook() == false {
		t.Error("GotoNextComicBook should return true when next comic book available")
	}
}

func TestGotPrevComicBookReturnsFalseWhenNoComicBooksLoaded(t *testing.T) {
	b := NewRoot()
	b.GotoLatestComicBook()
	if b.GotoPrevComicBook() == true {
		t.Error("GotoPrevComicBook should return false when no prev comic book available")
	}
}

func TestGotoPrevComicBookReturnsTrueWhenSomeComicBooksLoaded(t *testing.T) {
	b := NewRoot()
	b.LoadComicBooks(testComicBooks)
	b.GotoLatestComicBook()
	if b.GotoPrevComicBook() == false {
		t.Error("GotoPrevComicBook should return true when prev comic book available")
	}
}

func TestCycleForward(t *testing.T) {
	b := NewRoot()
	b.LoadComicBooks(testComicBooks)

	// starting with first book
	if *b.GetCurrentComicBook() != comicBook1 {
		t.Error("Starting comic book should be comicBook1")
	}

	// move to sercond book
	b.GotoNextComicBook()
	if *b.GetCurrentComicBook() != comicBook2 {
		t.Error("Second comic book should be comicBook2")
	}

	// move to third book
	b.GotoNextComicBook()
	if *b.GetCurrentComicBook() != comicBook3 {
		t.Error("Third comic book should be comicBook3")
	}
}

func TestCycleBackward(t *testing.T) {
	b := NewRoot()
	b.LoadComicBooks(testComicBooks)
	b.GotoLatestComicBook()

	// starting with third book
	if *b.GetCurrentComicBook() != comicBook3 {
		t.Error("Third comic book should be comicBook3")
	}

	// move to sercond book
	b.GotoPrevComicBook()
	if *b.GetCurrentComicBook() != comicBook2 {
		t.Error("Second comic book should be comicBook2")
	}

	// move to third first
	b.GotoPrevComicBook()
	if *b.GetCurrentComicBook() != comicBook1 {
		t.Error("First comic book should be comicBook1")
	}
}

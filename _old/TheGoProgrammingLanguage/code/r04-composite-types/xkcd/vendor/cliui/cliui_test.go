package cliui

import (
	"strings"
	"testing"
)

// Handler implements ui.Handler
type Handler struct {
	nPrevCommands  int
	nNextCommands  int
	lastGotoNumber uint
}

func NewHandler() *Handler {
	return &Handler{0, 0, 999}
}

func (h *Handler) OnPrevComicBook() {
	h.nPrevCommands++
}

func (h *Handler) OnNextComicBook() {
	h.nNextCommands++
}

func (h *Handler) OnGotoComicBook(number uint) {
	h.lastGotoNumber = number
}

func TestCliNext(t *testing.T) {
	reader := strings.NewReader(nextCmd)
	handler := NewHandler()
	cli := NewCliUI(reader, handler)
	cli.Run()

	if handler.nNextCommands != 1 {
		t.Errorf("Actual next commands != expected: %d != %d", handler.nNextCommands, 1)
	}
}

func TestCliPrev(t *testing.T) {
	reader := strings.NewReader(prevCmd)
	handler := NewHandler()
	cli := NewCliUI(reader, handler)
	cli.Run()

	if handler.nPrevCommands != 1 {
		t.Errorf("Actual prev commands != expected: %d != %d", handler.nPrevCommands, 1)
	}
}

func TestCliValidGoto(t *testing.T) {
	reader := strings.NewReader(gotoCmd + " 5")
	handler := NewHandler()
	cli := NewCliUI(reader, handler)
	cli.Run()

	if handler.lastGotoNumber != 4 { // 1-based numbering is translated to 0-based indexing
		t.Errorf("Actual goto number != expected: %d != %d", handler.lastGotoNumber, 5)
	}
}

func TestCliInvalidGoto(t *testing.T) {
	reader := strings.NewReader(gotoCmd) // invalid: missing comic book number
	handler := NewHandler()
	cli := NewCliUI(reader, handler)
	cli.Run()

	if handler.lastGotoNumber != 999 {
		t.Errorf("Actual goto number != expected: %d != %d", handler.lastGotoNumber, 999)
	}
}

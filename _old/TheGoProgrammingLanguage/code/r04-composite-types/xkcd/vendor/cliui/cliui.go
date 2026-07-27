package cliui

import (
	"bufio"
	"io"
	"sharedkernel"
	"strconv"
	"strings"
	"ui"

	tm "github.com/buger/goterm"
)

const prevCmd = "p"
const nextCmd = "n"
const gotoCmd = "g"
const helpCmd = "h"

// CliUI implements ui interface
type CliUI struct {
	reader    io.Reader
	handler   ui.Handler
	status    string
	comicBook *sharedkernel.ComicBook
}

// NewCliUI is CliUI constructor. reader provides string commands. handler handles parsed commands
func NewCliUI(reader io.Reader, handler ui.Handler) *CliUI {
	return &CliUI{reader, handler, "", &sharedkernel.ComicBook{}}
}

// SetUHHandler allows to set ui command handler
func (u *CliUI) SetUHHandler(handler ui.Handler) {
	u.handler = handler
}

// Run begins cli parsing loop
func (u *CliUI) Run() {
	scanner := bufio.NewScanner(u.reader)
	for scanner.Scan() {
		line := scanner.Text()
		u.parseLine(line)
	}
}

// DisplayStatus shows single line status
func (u *CliUI) DisplayStatus(status string) {
	u.status = status
	u.redrawScreen()
}

// DisplayComicBook shows single comic book
func (u *CliUI) DisplayComicBook(comicBook *sharedkernel.ComicBook) {
	u.comicBook = comicBook
	u.redrawScreen()
}

func (u *CliUI) parseLine(line string) {
	switch {
	case strings.HasPrefix(line, nextCmd):
		u.handler.OnNextComicBook()
	case strings.HasPrefix(line, prevCmd):
		u.handler.OnPrevComicBook()
	case strings.HasPrefix(line, gotoCmd):
		if number, ok := parseGotoCmd(line); ok {
			u.handler.OnGotoComicBook(number - 1) // comic books are indexed from 0 but numbered from 1
		} else {
			u.DisplayStatus("Invalid command: " + line)
		}

	default:
		u.DisplayStatus("n: next, p: prev, g 5: goto 5, h: help")
	}
}

func parseGotoCmd(line string) (uint, bool) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return 0, false
	}
	result, err := strconv.Atoi(parts[1])
	if err != nil || result < 0 {
		return 0, false
	}
	return uint(result), true
}

func (u *CliUI) redrawScreen() {
	tm.Clear()

	tm.MoveCursor(0, 2)
	tm.Println(u.status)
	tm.Println()
	tm.Println("=========================")
	tm.Printf("[%d] %s\n", u.comicBook.Number, u.comicBook.Title)
	tm.Println("-------------------------")
	tm.Println(u.comicBook.Transcription)
	tm.Println("=========================")
	tm.Println()
	tm.MoveCursor(0, 1)
	tm.Print("> ")
	tm.Flush()
}

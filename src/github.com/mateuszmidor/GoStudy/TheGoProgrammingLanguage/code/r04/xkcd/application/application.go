package main

import (
	"browser"
	browserport "browser/infrastructure"
	"cliui"
	"fmt"
	"jsonstorage"
	"os"
	"repository"
	"sharedkernel"
	"strconv"
	"ui"
	"xkcdfetcher"
)

// Application wires up components: CLI, Browser and Repository
// Implements: ui.Handler
type Application struct {
	cli              ui.UI
	browser          browserport.BrowserPort
	repository       *repository.Repository
	currentComicBook *sharedkernel.ComicBook
}

// NewApplication is Application constructor
func NewApplication() *Application {
	storage := jsonstorage.NewJSONStorage(jsonstorage.DefaultOpenForReading, jsonstorage.DefaultOpenForWriting)
	fetcher := xkcdfetcher.NewXkcdFetcher(xkcdfetcher.DefaultSingleComicBookFetcher)
	repository := repository.NewRepository(fetcher, storage)
	browser := browser.NewRoot()
	ui := cliui.NewCliUI(os.Stdin, nil)
	app := &Application{ui, browser, repository, nil}
	ui.SetUHHandler(app)

	return app
}

// Run loads comic books and starts processing CLI commands
func (a *Application) Run() {
	a.loadOfflineComicBooks()
	a.asyncUpdateComicBooks()
	a.cli.Run()
}

func (a *Application) loadOfflineComicBooks() {
	comicBooks, err := a.repository.LoadOfflineComicBooks()
	if err != nil {
		a.cli.DisplayStatus("Loading comic books failed")
	} else {
		a.cli.DisplayStatus("Loaded comic books")
		a.browser.LoadComicBooks(comicBooks)
		a.browseLatestIfNothingBrowsedYet()
	}
}

func (a *Application) browseLatestIfNothingBrowsedYet() {
	a.browser.GotoLatestComicBook()
	if a.currentComicBook == nil {
		a.browseCurrentComicBook()
	}
}
func (a *Application) browseCurrentComicBook() {
	a.currentComicBook = a.browser.GetCurrentComicBook()
	a.cli.DisplayComicBook(a.currentComicBook)
}

func (a *Application) asyncUpdateComicBooks() {
	a.cli.DisplayStatus("Updating comic books from xkcd.com...")
	_, err := a.repository.StartUpdateFromOnlineService(a.onComicBookFetched, a.onOfflineComicBooksUpdated)
	if err != nil {
		a.cli.DisplayStatus("Updating from xkcd.com failed")
	}
}

func (a *Application) onComicBookFetched(comicBookNo, total int) {
	status := fmt.Sprintf("Fetched %d/%d comic books", comicBookNo, total)
	a.cli.DisplayStatus(status)
}

func (a *Application) onOfflineComicBooksUpdated() {
	a.cli.DisplayStatus("Fetched all comic books. Reloading...")
	a.loadOfflineComicBooks()
}

// OnPrevComicBook implements ui.Handler interface
func (a *Application) OnPrevComicBook() {
	if a.browser.GotoPrevComicBook() {
		a.browseCurrentComicBook()
	} else {
		a.cli.DisplayStatus("No previous comic book")
	}
}

// OnNextComicBook implements ui.Handler interface
func (a *Application) OnNextComicBook() {
	if a.browser.GotoNextComicBook() {
		a.browseCurrentComicBook()
	} else {
		a.cli.DisplayStatus("No next comic book")
	}
}

// OnGotoComicBook implements ui.Handler interface
func (a *Application) OnGotoComicBook(number uint) {
	if a.browser.GotoComicBookNumber(number) {
		a.browseCurrentComicBook()
	} else {
		a.cli.DisplayStatus("No comic book number " + strconv.Itoa(int(number)))
	}
}

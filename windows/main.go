package main

import (
	"io"
	"net/http"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// prepare windowed application
	a := app.New()

	// prepare the main window
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(500, 100))

	// prepare a label to display text on it
	displayFactLabel := widget.NewLabel("Hello Fyne!")
	displayFactLabel.Wrapping = fyne.TextWrapBreak

	// prepare a button to fetch random facts from internet
	getFactButton := widget.NewButton("get random fact!", func() {
		displayFactLabel.SetText(getRandomFact())
	})

	// place label and button on the window
	w.SetContent(container.NewVBox(
		displayFactLabel,
		getFactButton,
	))

	// run the whole thing
	w.ShowAndRun()
}

// getRandomFact fetches a single random fact from internet
func getRandomFact() string {
	// get client wih timeout
	client := http.Client{Timeout: time.Second * 10}

	// prepare request
	req, err := http.NewRequest(http.MethodGet, "https://uselessfacts.jsph.pl/api/v2/facts/random", nil)
	if err != nil {
		return err.Error()
	}
	req.Header.Add("Accept", "text/plain")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	// read response
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	text := string(bytes)

	// return fact, without it's source footer
	return text[:strings.Index(text, "\n")]
}

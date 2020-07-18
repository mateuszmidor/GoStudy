package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mateuszmidor/GoStudy/flight-finder/cmd/util"
)

func main() {
	runWEB()
}

func runWEB() {
	finder := util.NewConnectionFinder("../../segments.csv.gz", "<br >")
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.HandleFunc("/api/find/text", handleFindAsText(finder))
	http.HandleFunc("/api/find/json", handleFindAsJSON(finder))
	http.ListenAndServe(":8080", nil)
}

func handleFindAsText(f *util.ConnectionFinder) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		from := strings.ToUpper(r.FormValue("from"))
		to := strings.ToUpper(r.FormValue("to"))

		w.Header().Set("Content-Type", "application/text")

		f.FindConnectionsAsText(w, from, to)

		fmt.Printf("%s -> %s\n", from, to)
	}
}

func handleFindAsJSON(f *util.ConnectionFinder) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		from := strings.ToUpper(r.FormValue("from"))
		to := strings.ToUpper(r.FormValue("to"))

		w.Header().Set("Content-Type", "application/json")

		f.FindConnectionsAsJSON(w, from, to)
		fmt.Printf("%s -> %s\n", from, to)
	}
}

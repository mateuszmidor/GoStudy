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
	http.HandleFunc("/api/find", handleFind(finder))
	http.ListenAndServe(":8080", nil)
}

func handleFind(f *util.ConnectionFinder) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		from := strings.ToUpper(r.FormValue("from"))
		to := strings.ToUpper(r.FormValue("to"))

		w.Header().Set("Content-Type", "application/text")

		f.FindConnections(w, from, to)

		fmt.Printf("%s -> %s\n", from, to)
	}
}

package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	finder := newWebPathFinder("../../segments.csv.gz")
	runWEB(finder)
}

func runWEB(f *webPathFinder) {
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.HandleFunc("/api/find", handleFind(f))
	http.ListenAndServe(":8080", nil)
}

func handleFind(f *webPathFinder) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		from := strings.ToUpper(r.FormValue("from"))
		to := strings.ToUpper(r.FormValue("to"))

		w.Header().Set("Content-Type", "application/text")

		start := time.Now()
		f.findConnections(from, to, w)
		d := time.Now().Sub(start)
		fmt.Fprintf(w, "Took %dms\n", d.Milliseconds())

		fmt.Printf("%s -> %s\n", from, to)
	}
}

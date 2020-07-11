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
	http.HandleFunc("/find", withFinder(f))
	http.ListenAndServe(":8080", nil)
}

func withFinder(f *webPathFinder) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		from := strings.ToUpper(r.URL.Query().Get("from"))
		to := strings.ToUpper(r.URL.Query().Get("to"))

		w.Header().Set("Content-Type", "text/html")

		start := time.Now()
		f.findConnections(from, to, w)
		d := time.Now().Sub(start)
		fmt.Fprintf(w, "Took %dms\n", d.Milliseconds())

		fmt.Printf("%s -> %s\n", from, to)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/cmd/util"
)

func main() {
	runWEB()
}

func runWEB() {
	finder := util.NewConnectionFinder("../../../data/segments.csv.gz", "../../../data/airports.csv.gz", "../../../data/nations.csv.gz", "<br >")
	http.Handle("/", &templateHandler{filename: "map.html"})
	http.Handle("/list", &templateHandler{filename: "list.html"})
	http.HandleFunc("/api/find/text", handleFindAsText(finder))
	http.HandleFunc("/api/find/json", handleFindAsJSON(finder))
	http.ListenAndServe(":9000", nil)
}

func handleFindAsText(f *util.ConnectionFinder) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		from := strings.ToUpper(r.FormValue("from"))
		to := strings.ToUpper(r.FormValue("to"))
		maxSegmentCount, _ := strconv.Atoi(r.FormValue("maxsegmentcount"))

		w.Header().Set("Content-Type", "application/text")

		f.FindConnectionsAsText(w, from, to, maxSegmentCount)

		fmt.Printf("%s -> %s\n", from, to)
	}
}

func handleFindAsJSON(f *util.ConnectionFinder) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		from := strings.ToUpper(r.FormValue("from"))
		to := strings.ToUpper(r.FormValue("to"))
		maxSegmentCount, _ := strconv.Atoi(r.FormValue("maxsegmentcount"))
		if err := f.FindConnectionsAsJSON(w, from, to, maxSegmentCount); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			dumpErrorAsJSON(w, err)
		}

		fmt.Printf("%s -> %s\n", from, to)
	}
}

func dumpErrorAsJSON(w io.Writer, e error) {
	errorView := struct {
		Error string `json:"error"`
	}{
		e.Error(),
	}

	json.NewEncoder(w).Encode(errorView)
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/application"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/csv"
)

func main() {
	runWEB()
}

func runWEB() {
	repo := csv.NewFlightsDataRepoCSV("../../../data/")
	finder := application.NewConnectionFindingService(repo)

	http.Handle("/", &templateHandler{filename: "map.html"})
	http.HandleFunc("/api/find/json", handleFindAsJSON(finder))
	http.ListenAndServe(":9000", nil)
}

func handleFindAsJSON(f *application.ConnectionFindingService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		from := strings.ToUpper(r.FormValue("from"))
		to := strings.ToUpper(r.FormValue("to"))
		maxSegmentCount, _ := strconv.Atoi(r.FormValue("maxsegmentcount"))
		renderer := application.NewPathRendererAsJSON(w)
		if err := f.Find(from, to, maxSegmentCount, renderer); err != nil {
			fmt.Printf("%v\n", err)
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

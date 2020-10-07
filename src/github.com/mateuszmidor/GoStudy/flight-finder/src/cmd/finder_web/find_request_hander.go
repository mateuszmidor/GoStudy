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

type findRequestHandler struct {
	finder *application.ConnectionFinder
}

func NewFindRequestHandler() *findRequestHandler {
	repo := csv.NewFlightsDataRepoCSV("../../../data/")
	finder := application.NewConnectionFinder(repo)
	return &findRequestHandler{finder: finder}
}

func (h *findRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	from := strings.ToUpper(r.FormValue("from"))
	to := strings.ToUpper(r.FormValue("to"))
	maxSegmentCount, _ := strconv.Atoi(r.FormValue("maxsegmentcount"))
	renderer := application.NewPathRendererAsJSON(w)
	if err := h.finder.Find(from, to, maxSegmentCount, renderer); err != nil {
		fmt.Printf("%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		dumpErrorAsJSON(w, err)
	}

	fmt.Printf("%s -> %s\n", from, to)
}

func dumpErrorAsJSON(w io.Writer, e error) {
	errorView := struct {
		Error string `json:"error"`
	}{
		e.Error(),
	}

	json.NewEncoder(w).Encode(errorView)
}

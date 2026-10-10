// Package power provides the /power endpoint: a raised to the b-th power.
package power

import (
	"math"
	"net/http"

	"go-arch-lint/api"
)

// Register mounts the "GET /power" handler on the given mux.
func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /power", powerHandler)
}

func powerHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	base, err := api.ParseParam(q, "a")
	if err != nil {
		api.ReturnError(w, http.StatusBadRequest, err.Error())
		return
	}
	exponent, err := api.ParseParam(q, "b")
	if err != nil {
		api.ReturnError(w, http.StatusBadRequest, err.Error())
		return
	}

	result := math.Pow(base, exponent)
	if math.IsNaN(result) || math.IsInf(result, 0) {
		api.ReturnError(w, http.StatusBadRequest, "power result is not a finite number")
		return
	}

	api.ReturnSuccess(w, result)
}

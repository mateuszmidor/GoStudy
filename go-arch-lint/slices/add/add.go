// Package add provides the /add endpoint: the sum of two numbers.
package add

import (
	"net/http"

	"go-arch-lint/api"
)

// Register mounts the "GET /add" handler on the given mux.
func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /add", addHandler)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	a, err := api.ParseParam(q, "a")
	if err != nil {
		api.ReturnError(w, http.StatusBadRequest, err.Error())
		return
	}
	b, err := api.ParseParam(q, "b")
	if err != nil {
		api.ReturnError(w, http.StatusBadRequest, err.Error())
		return
	}

	api.ReturnSuccess(w, a+b)
}

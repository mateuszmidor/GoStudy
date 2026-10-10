// Package multiply provides the /multiply endpoint: the product of two numbers.
package multiply

import (
	"net/http"

	"go-arch-lint/api"
)

// Register mounts the "GET /multiply" handler on the given mux.
func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /multiply", multiplyHandler)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
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

	api.ReturnSuccess(w, a*b)
}

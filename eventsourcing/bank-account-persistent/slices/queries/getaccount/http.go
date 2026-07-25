package getaccount

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type HTTPHandler struct {
	handler *QueryHandler
}

func NewHTTPHandler(handler *QueryHandler) *HTTPHandler {
	return &HTTPHandler{		handler: handler}
}

func (h* HTTPHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /accounts/{id}", h.Handle)
}

func (h *HTTPHandler) Handle (w http.ResponseWriter, r* http.Request) {
	// get http request parameters
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// do the work
	result, err := h.handler.HandleQuery(r.Context(), GetAccount{AccountID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// return http result
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
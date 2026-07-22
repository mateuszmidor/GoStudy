package listaccounts

import (
	"encoding/json"
	"net/http"
)

type HTTPHandler struct {
	handler *QueryHandler
}

func NewHTTPHandler(handler *QueryHandler) *HTTPHandler {
	return &HTTPHandler{handler: handler}
}

func (h *HTTPHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /accounts", h.Handle)
}

func (h *HTTPHandler) Handle(w http.ResponseWriter, req *http.Request) {
	result, err := h.handler.HandleQuery(req.Context(), ListAccounts{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

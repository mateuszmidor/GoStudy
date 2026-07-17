package listaccounts

import (
	"encoding/json"
	"log/slog"
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
	slog.Info(req.Method + " " + req.URL.Path)

	result, err := h.handler.HandleQuery(req.Context(), ListAccounts{})
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

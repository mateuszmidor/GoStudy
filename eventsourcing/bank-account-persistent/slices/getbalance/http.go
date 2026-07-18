package getbalance

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type HTTPHandler struct {
	handler *QueryHandler
}

func NewHTTPHandler(handler *QueryHandler) *HTTPHandler {
	return &HTTPHandler{handler: handler}
}

func (h *HTTPHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /accounts/{id}/balance", h.Handle)
}

func (h *HTTPHandler) Handle(w http.ResponseWriter, req *http.Request) {
	slog.Info(req.Method + " " + req.URL.Path)

	id, err := uuid.Parse(req.PathValue("id"))
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.handler.HandleQuery(req.Context(), GetBalance{AccountID: id})
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

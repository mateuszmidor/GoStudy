package fundaccount

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/terraskye/eventsourcing"
)

type request struct {
	Dollars uint
}

type HTTPHandler struct {
	handler eventsourcing.CommandHandler[FundAccount]
}

func NewHTTPHandler(handler eventsourcing.CommandHandler[FundAccount]) *HTTPHandler {
	return &HTTPHandler{handler: handler}
}

func (h *HTTPHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /accounts/{id}/atm", h.Handle)
}

func (h *HTTPHandler) Handle(w http.ResponseWriter, req *http.Request) {
	var r request
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slog.Info(req.Method+" "+req.URL.Path, slog.Any("payload", r))
	id, err := uuid.Parse(req.PathValue("id"))
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := FundAccount{
		AccountID: id,
		Dollars:   r.Dollars,
	}
	if _, err := h.handler(req.Context(), cmd); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

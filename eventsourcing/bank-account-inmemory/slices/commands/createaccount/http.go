package createaccount

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/terraskye/eventsourcing"
)

type request struct {
	OwnerName string `json:"owner_name"`
}

type HTTPHandler struct {
	handler eventsourcing.CommandHandler[CreateAccount]
}

func NewHTTPHandler(handler eventsourcing.CommandHandler[CreateAccount]) *HTTPHandler {
	return &HTTPHandler{handler: handler}
}

func (h *HTTPHandler) Register(m *http.ServeMux) {
	m.HandleFunc("POST /accounts", h.Handle)
}

func (h *HTTPHandler) Handle(w http.ResponseWriter, req *http.Request) {
	var r request
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slog.Info(req.Method+" "+req.URL.Path, slog.Any("payload", r))

	cmd := CreateAccount{
		AccountID: uuid.New(),
		OwnerName: r.OwnerName,
	}
	if _, err := h.handler(req.Context(), cmd); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rsp := map[string]any{
		"account_id": cmd.AccountID,
	}
	json.NewEncoder(w).Encode(rsp)
}

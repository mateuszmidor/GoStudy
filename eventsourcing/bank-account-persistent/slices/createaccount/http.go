package createaccount

import (
	"encoding/json"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := CreateAccount{
		AccountID: uuid.New(),
		OwnerName: r.OwnerName,
	}
	if _, err := h.handler(req.Context(), cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rsp := map[string]any{
		"account_id": cmd.AccountID,
	}
	json.NewEncoder(w).Encode(rsp)
}

package presentation

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"bank-account/application"
	"bank-account/infrastructure"
)

// HttpController exposes the account use cases over HTTP.
type HttpController struct {
	repository *infrastructure.Repository
}

// NewHttpController creates a controller wired to the repository.
func NewHttpController(repository *infrastructure.Repository) *HttpController {
	return &HttpController{
		repository: repository,
	}
}

// RegisterRoutes connects the controller handlers to HTTP routes.
func (c *HttpController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /accounts", c.handleCreateAccount)
	mux.HandleFunc("POST /accounts/{id}/deposits", c.handleDeposit)
	mux.HandleFunc("GET /accounts/{id}", c.handleGetAccount)
	mux.HandleFunc("GET /accounts", c.handleListAccounts)
}

// handleCreateAccount creates a new account aggregate and persists its initial event.
func (c *HttpController) handleCreateAccount(w http.ResponseWriter, req *http.Request) {
	// HTTP preamble
	var payload struct {
		OwnerName string `json:"owner_name"`
	}
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Info(req.Method+" "+req.URL.Path, slog.Any("payload", payload))

	// Call business logic
	account, err := application.NewAccount(payload.OwnerName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.repository.Save(account.FlushEvents()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HTTP Postamble
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{"account_id": account.ID})
}

// handleDeposit loads the target account, applies a deposit, and saves the new event.
func (c *HttpController) handleDeposit(w http.ResponseWriter, req *http.Request) {
	// HTTP preamble
	accountID, err := uuid.Parse(req.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload struct {
		Dollars uint `json:"dollars"`
	}
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Info(req.Method+" "+req.URL.Path, slog.Any("payload", payload), slog.String("account_id", accountID.String()))

	// Call business logic
	account, err := c.repository.Get(accountID)
	if err != nil {
		if errors.Is(err, infrastructure.ErrAccountNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := account.Deposit(payload.Dollars); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.repository.Save(account.FlushEvents()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HTTP Postamble
	w.WriteHeader(http.StatusOK)
}

// handleGetAccount returns the current reconstructed account state as JSON.
func (c *HttpController) handleGetAccount(w http.ResponseWriter, req *http.Request) {
	// HTTP preamble
	accountID, err := uuid.Parse(req.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Info(req.Method+" "+req.URL.Path, slog.Any("payload", accountID.String()))

	// Call business logic
	account, err := c.repository.Get(accountID)
	if err != nil {
		if errors.Is(err, infrastructure.ErrAccountNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HTTP Postamble
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"account_id": account.ID,
		"owner_name": account.OwnerName,
		"balance":    account.Balance,
	})
}

// handleListAccounts returns every reconstructed account as a JSON array.
func (c *HttpController) handleListAccounts(w http.ResponseWriter, req *http.Request) {
	// HTTP preamble
	slog.Info(req.Method+" "+req.URL.Path)

	// Call business logic
	accounts, err := c.repository.List()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := make([]map[string]any, 0, len(accounts))
	for _, account := range accounts {
		result = append(result, map[string]any{
			"account_id": account.ID,
			"owner_name": account.OwnerName,
			"balance":    account.Balance,
		})
	}

	// HTTP Postamble
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

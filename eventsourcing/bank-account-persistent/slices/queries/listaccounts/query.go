package listaccounts

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ListAccounts struct{}

func (q ListAccounts) ID() []byte { return []byte("list-accounts") }

type Account struct {
	ID uuid.UUID
	// Dollars   uint // for account balance use GetBalance query
	OwnerName string
	CreatedAt time.Time
}

type QueryHandler struct {
	projector *Projector
}

func NewQueryHandler(projector *Projector) *QueryHandler {
	return &QueryHandler{projector: projector}
}

// HandleQuery returns the cached account list
func (h *QueryHandler) HandleQuery(ctx context.Context, _ ListAccounts) ([]Account, error) {
	// return accounts from cache
	return h.projector.GetAll(), nil
}

package application

import "github.com/google/uuid"

// AccountList holds list of accounts.
type AccountList []Account

// GetByID returns the account with the matching ID, or nil if it is missing.
func (list AccountList) GetByID(accountID uuid.UUID) *Account {
	for account := range list {
		if list[account].ID == accountID {
			return &list[account]
		}
	}
	return nil
}
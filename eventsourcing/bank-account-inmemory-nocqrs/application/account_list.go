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

func (list* AccountList) GetOrInsert(accountID uuid.UUID) *Account {
	account := list.GetByID(accountID)
	if account != nil {
		return account
	} else {
		account := Account{ID: accountID}
		*list = append(*list, account)
		return  &(*list)[len(*list)-1] // return pointer to the newly added account
	}
}
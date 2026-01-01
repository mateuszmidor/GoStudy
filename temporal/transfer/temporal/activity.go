package temporal

import (
	"context"
	"fmt"
)

var counter = 1

func Withdraw(ctx context.Context, account string, amountPLN int) error {
	// check invariants
	if account == "" {
		return fmt.Errorf("withdrawal-missing-account-error")
	}
	if amountPLN <= 0 {
		return fmt.Errorf("withdrawal-wrong-amount-error")
	}

	// success
	return nil
}

func Deposit(ctx context.Context, account string, amountPLN int) error {
	// simulate a single error
	if counter > 0 {
		counter--
		return fmt.Errorf("withdrawal-error")
	}

	// check invariants
	if account == "" {
		return fmt.Errorf("deposit-missing-account-error")
	}
	if amountPLN <= 0 {
		return fmt.Errorf("deposit-wrong-amount-error")
	}

	// success
	return nil
}

package temporal

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/temporal"
)

var counter = 2

func Withdraw(ctx context.Context, account string, amountPLN int) error {
	// check invariants
	if account == "" {
		return temporal.NewApplicationError("missing account", "MissingAccountError")
	}
	if amountPLN <= 0 {
		return temporal.NewApplicationError("wrong amount", "WrongAmountError")
	}

	// success
	return nil
}

func Deposit(ctx context.Context, account string, amountPLN int) error {
	// simulate a deposit error
	if counter > 0 {
		counter--
		return fmt.Errorf("deposit-error")
	}

	// check invariants
	if account == "" {
		return temporal.NewApplicationError("missing account", "MissingAccountError")
	}
	if amountPLN <= 0 {
		return temporal.NewApplicationError("wrong amount", "WrongAmountError")
	}

	// success
	return nil
}

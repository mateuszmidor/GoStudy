package temporal

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type TransferDetails struct {
	AccountFrom string `json:"accountFrom"`
	AccountTo   string `json:"accountTo"`
	AmountPLN   int    `json:"amountPLN"`
}

func MoneyTransferWorkflow(ctx workflow.Context, details TransferDetails) (string, error) {
	rp := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        500, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"MissingAccountError", "WrongAmountError"},
	}
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy:         rp,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	if err := workflow.ExecuteActivity(ctx, Withdraw, details.AccountFrom, details.AmountPLN).Get(ctx, nil); err != nil {
		return "", err
	}
	if err := workflow.ExecuteActivity(ctx, Deposit, details.AccountTo, details.AmountPLN).Get(ctx, nil); err != nil {
		return "", err
	}

	return fmt.Sprintf("transferred %d PLN %s->%s", details.AmountPLN, details.AccountFrom, details.AccountTo), nil
}

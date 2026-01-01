package temporal

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func MoneyTransferWorkflow(ctx workflow.Context, accountFrom, accountTo string, amountPLN int) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	if err := workflow.ExecuteActivity(ctx, Withdraw, accountFrom, amountPLN).Get(ctx, nil); err != nil {
		return "", err
	}
	if err := workflow.ExecuteActivity(ctx, Deposit, accountTo, amountPLN).Get(ctx, nil); err != nil {
		return "", err
	}

	return fmt.Sprintf("transferred %d PLN %s->%s", amountPLN, accountFrom, accountTo), nil
}

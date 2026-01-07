package main

import (
	"log"
	"time"

	"go.temporal.io/sdk/workflow"
)

func WaitForSignalWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// wait for signal
	log.Println("waiting for signal continue_workflow...")
	continuationSelector := workflow.NewSelector(ctx)
	continuationSignalChan := workflow.GetSignalChannel(ctx, "continue_workflow")
	continuationSelector.AddReceive(continuationSignalChan, func(c workflow.ReceiveChannel, more bool) {
		var msg string
		c.Receive(ctx, &msg)
		log.Println("received signal continue_workflow:", msg)
	})
	continuationSelector.Select(ctx) // here we wait for signal to arrive

	// run activity
	var result string
	err := workflow.ExecuteActivity(ctx, Greet, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	// SUCCESS
	return result, nil
}

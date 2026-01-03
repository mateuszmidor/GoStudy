package main

import (
	"context"
	"log"

	"github.com/mateuszmidor/GoStudy/temporal/moneytransfer/temporal"
	"go.temporal.io/sdk/client"
)

func main() {
	startWorkflow()
}

func startWorkflow() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "money-transfer-workflow",
		TaskQueue: "my-task-queue",
	}

	details := temporal.TransferDetails{AccountFrom: "ING", AccountTo: "BNP", AmountPLN: 325}
	we, err := c.ExecuteWorkflow(context.Background(), options, temporal.MoneyTransferWorkflow, details)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}

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

	we, err := c.ExecuteWorkflow(context.Background(), options, temporal.MoneyTransferWorkflow, "ING", "BNP", 325)
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

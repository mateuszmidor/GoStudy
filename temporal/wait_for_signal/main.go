package main

import (
	"context"
	"log"
	"time"

	"go.temporal.io/sdk/client"
)

func main() {
	go startWorker()
	startWorkflow()
}

func startWorkflow() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "wait-for-signal-workflow",
		TaskQueue: "my-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, WaitForSignalWorkflow, "Mateusz")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// signal the workflow after a delay
	time.Sleep(3 * time.Second)
	if err := c.SignalWorkflow(context.Background(), we.GetID(), we.GetRunID(), "continue_workflow", "go now!"); err != nil {
		log.Fatalf("failed to signal workflow: %+v", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}

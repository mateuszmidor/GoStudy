package main

import (
	"log"

	"github.com/mateuszmidor/GoStudy/temporal/moneytransfer/temporal"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "my-task-queue", worker.Options{})

	w.RegisterWorkflow(temporal.MoneyTransferWorkflow)
	w.RegisterActivity(temporal.Deposit)
	w.RegisterActivity(temporal.Withdraw)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

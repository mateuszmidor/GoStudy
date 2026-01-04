package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/mateuszmidor/GoStudy/temporal/geolocation"
	"go.temporal.io/sdk/client"
)

func main() {
	if len(os.Args) <= 1 {
		os.Args = append(os.Args, "Andrzej")
	}
	name := os.Args[1]

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowID := "getAddressFromIP-" + uuid.New().String()

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: geolocation.TaskQueueName,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, geolocation.GetAddressFromIP, name)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}

	fmt.Println(result)
}

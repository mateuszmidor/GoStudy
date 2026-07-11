// main.go
package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	membus "github.com/terraskye/eventsourcing/eventbus/memory"
	memstore "github.com/terraskye/eventsourcing/eventstore/memory"

	"task-management/slices/completetask"
	"task-management/slices/createtask"
	"task-management/slices/tasklist"
)

func main() {
	store := memstore.NewMemoryStore(100)
	defer store.Close()

	bus := membus.NewEventBus(100)
	defer bus.Close()

	// Create the projector and register it on the bus.
	projector := tasklist.NewProjector()
	if err := bus.Subscribe(context.Background(), "task-list-projector", projector.EventHandlers()); err != nil {
		log.Fatal(err)
	}

	// Forward events from the store to the bus.
	go func() {
		for env := range store.Events() {
			bus.Dispatch(env)
		}
	}()

	// Handlers
	createTaskHTTP := createtask.NewHTTPHandler(createtask.NewHandler(store))
	completeTaskHTTP := completetask.NewHTTPHandler(completetask.NewHandler(store))
	listTasksHTTP := tasklist.NewHTTPHandler(tasklist.NewQueryHandler(projector))

	r := gin.Default()
	tasks := r.Group("/api/v1/tasks")
	createTaskHTTP.RegisterRoutes(tasks)
	completeTaskHTTP.RegisterRoutes(tasks)
	listTasksHTTP.RegisterRoutes(tasks)

	log.Fatal(r.Run(":8080"))
}

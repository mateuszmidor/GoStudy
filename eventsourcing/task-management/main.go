// main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/terraskye/eventsourcing/eventstore/memory"

	"task-management/slices/createtask"
)

func main() {
	store := memory.NewMemoryStore(100)
	defer store.Close()

	createTaskHandler := createtask.NewHandler(store)
	createTaskHTTP := createtask.NewHTTPHandler(createTaskHandler)

	r := gin.Default()
	tasks := r.Group("/api/v1/tasks")
	createTaskHTTP.RegisterRoutes(tasks)

	log.Fatal(r.Run(":8080"))
}

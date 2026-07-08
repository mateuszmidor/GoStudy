// main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/terraskye/eventsourcing/eventstore/memory"

	"task-management/slices/completetask"
	"task-management/slices/createtask"
	"task-management/slices/listtasks"
)

func main() {
	store := memory.NewMemoryStore(100)
	defer store.Close()

	createTaskHandler := createtask.NewHandler(store)
	createTaskHTTP := createtask.NewHTTPHandler(createTaskHandler)

	listTasksHandler := listtasks.NewQueryHandler(store)
	listTasksHTTP := listtasks.NewHTTPHandler(listTasksHandler)

	completeTaskHandler := completetask.NewHandler(store)
	completeTaskHTTP := completetask.NewHTTPHandler(completeTaskHandler)

	r := gin.Default()
	tasks := r.Group("/api/v1/tasks")
	createTaskHTTP.RegisterRoutes(tasks)
	listTasksHTTP.RegisterRoutes(tasks) // same router group as createtask
	completeTaskHTTP.RegisterRoutes(tasks)

	log.Fatal(r.Run(":8080"))
}

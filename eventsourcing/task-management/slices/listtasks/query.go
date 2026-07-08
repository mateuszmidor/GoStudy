// slices/listtasks/query.go
package listtasks

import (
	"context"
	"errors"
	"fmt"

	"task-management/events"

	"github.com/terraskye/eventsourcing"
)

// ListTasks is the query type. ID() is used as a cache key (empty = no specific aggregate).
type ListTasks struct{}

func (q ListTasks) ID() []byte { return []byte("list-tasks") }

// Task is one item in the read model.
type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// TaskList is the read model returned by the query.
type TaskList struct {
	Tasks []Task `json:"tasks"`
}

// evolve builds the TaskList from a stream of events.
func evolve(state *TaskList, envelope *eventsourcing.Envelope) *TaskList {
	switch e := envelope.Event.(type) {
	case *events.TaskCreated:
		state.Tasks = append(state.Tasks, Task{
			ID:    e.TaskID.String(),
			Title: e.Title,
		})
	}
	return state
}

// QueryHandler rebuilds the read model on every call.
type QueryHandler struct {
	store eventsourcing.EventStore
}

func NewQueryHandler(store eventsourcing.EventStore) *QueryHandler {
	return &QueryHandler{store: store}
}

func (h *QueryHandler) HandleQuery(ctx context.Context, _ ListTasks) (*TaskList, error) {
	iter, err := h.store.LoadFromAll(ctx, eventsourcing.Revision(0))
	if err != nil {
		if errors.Is(err, eventsourcing.ErrInvalidRevision) {
			return &TaskList{Tasks: make([]Task, 0)}, nil
		}
		return nil, fmt.Errorf("list tasks: %w", err)
	}

	result := &TaskList{Tasks: make([]Task, 0)}
	for iter.Next(ctx) {
		result = evolve(result, iter.Value())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

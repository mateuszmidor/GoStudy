// slices/tasklist/projector.go
package tasklist

import (
	"context"
	"maps"
	"slices"
	"sync"

	"task-management/events"

	"github.com/terraskye/eventsourcing"
)

// Task is the cached read model entry.
type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Archived  bool   `json:"archived"`
}

// TaskList is the cached read model served by queries.
type TaskList struct {
	Tasks []Task
}

// Projector maintains the cached TaskList.
type Projector struct {
	mu    sync.RWMutex
	tasks map[string]*Task // key is Task.ID
}

func NewProjector() *Projector {
	return &Projector{tasks: make(map[string]*Task)}
}

// All returns a snapshot of the current task list.
func (p *Projector) All() []*Task {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return slices.Collect(maps.Values(p.tasks))
}

// OnTaskCreated is called by the event bus when a TaskCreated event arrives.
func (p *Projector) OnTaskCreated(_ context.Context, e *events.TaskCreated) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.tasks[e.TaskID.String()] = &Task{
		ID:    e.TaskID.String(),
		Title: e.Title,
	}
	return nil
}

// OnTaskCompleted is called by the event bus when a TaskCompleted event arrives.
func (p *Projector) OnTaskCompleted(_ context.Context, e *events.TaskCompleted) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if t, ok := p.tasks[e.TaskID.String()]; ok {
		t.Completed = true
	}
	return nil
}

// OnTaskArchived is called by the event bus when a TaskArchived event arrives.
func (p *Projector) OnTaskArchived(_ context.Context, e *events.TaskArchived) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if t, ok := p.tasks[e.TaskID.String()]; ok {
		t.Archived = true
	}
	return nil
}

// EventHandlers returns the typed handlers to register on the event bus.
func (p *Projector) EventHandlers() *eventsourcing.EventGroupProcessor {
	return eventsourcing.NewEventGroupProcessor(
		eventsourcing.OnEvent(p.OnTaskCreated),
		eventsourcing.OnEvent(p.OnTaskCompleted),
		eventsourcing.OnEvent(p.OnTaskArchived),
	)
}

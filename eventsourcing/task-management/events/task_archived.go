// events/task_archived.go
package events

import (
	"time"

	"github.com/google/uuid"
)

type TaskArchived struct {
	TaskID     uuid.UUID `json:"task_id"`
	ArchivedAt time.Time `json:"archived_at"`
}

func (e *TaskArchived) AggregateID() string { return e.TaskID.String() }
func (e *TaskArchived) EventType() string   { return "TaskArchived" }

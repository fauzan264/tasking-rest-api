package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID
	Task      string
	Assign    string
	Status    int
	Deadline  string
	CreatedAt time.Time
}

type TaskStatus struct {
	Id     uuid.UUID
	Status int
}

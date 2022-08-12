package task

import (
	"github.com/google/uuid"
)

type TaskFormatter struct {
	Id       uuid.UUID `json:"id"`
	Task     string    `json:"task"`
	Assign   string    `json:"assign"`
	Status   int       `json:"status"`
	Deadline string    `json:"deadline"`
}

func FormatTask(task Task) TaskFormatter {
	taskFormatter := TaskFormatter{}
	taskFormatter.Id = task.Id
	taskFormatter.Task = task.Task
	taskFormatter.Assign = task.Assign
	taskFormatter.Status = task.Status
	taskFormatter.Deadline = task.Deadline

	return taskFormatter
}

func FormatTasks(tasks []Task) []TaskFormatter {
	tasksFormatter := []TaskFormatter{}

	for _, task := range tasks {
		taskFormatter := FormatTask(task)
		tasksFormatter = append(tasksFormatter, taskFormatter)
	}

	return tasksFormatter
}

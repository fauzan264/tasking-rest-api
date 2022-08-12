package task

import (
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GetTasks(Id string) ([]Task, error)
	CreateTask(input CreateTaskInput) (Task, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTasks(Id string) ([]Task, error) {

	if Id != "" {
		tasks, err := s.repository.FindByTaskID(uuid.MustParse(Id))
		if err != nil {
			return tasks, err
		}

		return tasks, nil
	}

	tasks, err := s.repository.FindAll()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *service) CreateTask(input CreateTaskInput) (Task, error) {
	task := Task{}
	task.Task = input.Task
	task.Assign = input.Assign

	// parse string date to golang time
	_, err := time.Parse("01/02/2006", input.Deadline)
	if err != err {
		return task, err
	}
	task.Deadline = input.Deadline

	id, err := uuid.NewRandom()
	if err != nil {
		return task, err
	}
	task.Id = id

	newTask, err := s.repository.Save(task)
	if err != nil {
		return newTask, err
	}

	return newTask, nil
}

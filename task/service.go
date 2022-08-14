package task

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GetTasks(Id string) ([]Task, error)
	GetTaskById(input GetTaskDetailInput) (Task, error)
	CreateTask(input CreateTaskInput) (Task, error)
	UpdateData(Id GetTaskDetailInput, InputData CreateTaskInput) (Task, error)
	UpdateDataStatus(Id GetTaskDetailInput, InputData CreateTaskStatusInput) (Task, error)
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

func (s *service) GetTaskById(input GetTaskDetailInput) (Task, error) {

	task, err := s.repository.FindById(uuid.MustParse(input.Id))

	if err != nil {
		return task, err
	}

	return task, nil
}

func (s *service) UpdateData(Id GetTaskDetailInput, InputData CreateTaskInput) (Task, error) {
	task, err := s.repository.FindById(uuid.MustParse(Id.Id))

	if err != nil {
		return task, err
	}

	task.Task = InputData.Task
	task.Assign = InputData.Assign

	// deadline, err := time.Parse("2006-01-02T15:04:05.000Z", InputData.Deadline)
	deadline, err := time.Parse("2006-01-02", InputData.Deadline)
	if err != nil {
		return task, err
	}
	task.Deadline = deadline.Format("2006-01-02")

	fmt.Println("tanggal deadline nya adalah", task.Deadline)
	UpdateTask, err := s.repository.UpdateData(task)
	if err != nil {
		return UpdateTask, err
	}

	return UpdateTask, err
}

func (s *service) UpdateDataStatus(Id GetTaskDetailInput, InputData CreateTaskStatusInput) (Task, error) {
	task, err := s.repository.FindById(uuid.MustParse(Id.Id))
	if err != nil {
		return task, err
	}

	task.Status = InputData.Status
	updateTaskStatus, err := s.repository.UpdateDataStatus(task)
	if err != nil {
		return updateTaskStatus, err
	}

	return updateTaskStatus, nil
}

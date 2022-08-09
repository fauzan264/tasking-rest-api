package task

import "github.com/google/uuid"

type Service interface {
	FindTasks(Id uuid.UUID) ([]Task, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindTasks(Id uuid.UUID) ([]Task, error) {
	if Id != uuid.Nil {
		tasks, err := s.repository.FindByTaskID(Id)
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

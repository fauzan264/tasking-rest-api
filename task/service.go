package task

import "github.com/google/uuid"

type Service interface {
	GetTasks(Id string) ([]Task, error)
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

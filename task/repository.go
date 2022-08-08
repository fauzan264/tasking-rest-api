package task

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Task, error)
	FindByTaskID(Id uuid.UUID) ([]Task, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (r *repository) FindByTaskID(Id uuid.UUID) ([]Task, error) {
	var tasks []Task

	err := r.db.Where("id = ?", Id).Find(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil

}

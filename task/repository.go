package task

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Task, error)
	FindByTaskID(Id uuid.UUID) ([]Task, error)
	FindById(Id uuid.UUID) (Task, error)
	Save(task Task) (Task, error)
	UpdateData(task Task) (Task, error)
	UpdateDataStatus(task Task) (Task, error)
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

func (r *repository) Save(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) FindById(Id uuid.UUID) (Task, error) {
	var task Task

	err := r.db.Where("id = ?", Id).Find(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) UpdateData(task Task) (Task, error) {
	err := r.db.Save(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) UpdateDataStatus(task Task) (Task, error) {
	err := r.db.Model(&task).Where("id = ?", task.Id).Update("status", task.Status).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

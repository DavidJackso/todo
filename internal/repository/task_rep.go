package repository

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TaskRepositoryGorm struct {
	db *gorm.DB
}

func NewTaskRepositoryGorm(db *gorm.DB) *TaskRepositoryGorm {
	return &TaskRepositoryGorm{
		db: db,
	}
}

func (r *TaskRepositoryGorm) CreateTask(task models.Task) (int, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		logrus.Error("Error add new task")
		return 0, nil
	}
	return int(task.ID), nil
}

func (r *TaskRepositoryGorm) DeleteTask(id int) error {
	return nil
}

func (r *TaskRepositoryGorm) UpdateTask(id int) (models.Task, error) {
	return models.Task{}, nil
}

func (r *TaskRepositoryGorm) GetTask(id int) (models.Task, error) {
	return models.Task{}, nil
}

func (r *TaskRepositoryGorm) GetAllTasks() ([]models.Task, error) {
	return nil, nil
}

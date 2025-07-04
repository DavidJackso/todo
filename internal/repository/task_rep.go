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
	task, err := getByID(id, *r.db)
	if err != nil {
		logrus.Error("bad")
	}
	result := r.db.Delete(&task)
	if result.Error != nil {
		logrus.Error(err)
	}
	return nil
}

func (r *TaskRepositoryGorm) UpdateTask(id int) (models.Task, error) {
	return models.Task{}, nil
}

func (r *TaskRepositoryGorm) GetTask(id int) (models.Task, error) {
	task, err := getByID(id, *r.db)
	if err.Error != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepositoryGorm) GetTasks() ([]models.Task, error) {
	return nil, nil
}
func getByID(id int, db gorm.DB) (models.Task, error) {
	var task models.Task
	result := db.Where("id = ?", id).First(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}

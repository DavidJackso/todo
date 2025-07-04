package repository

import (
	"fmt"

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

func (r *TaskRepositoryGorm) CreateTask(task models.Task, userID int) (int, error) {
	r.CreateCategory("aba")
	fmt.Print(task.CategoryID)
	task.UserID = uint(userID)
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

func (r *TaskRepositoryGorm) UpdateTask(id int, updTask models.Task) (models.Task, error) {
	task, err := getByID(id, *r.db)
	if err != nil {
		return models.Task{}, err
	}
	updTask.ID = task.ID

	task = updTask

	r.db.Save(task)

	return task, nil
}

func (r *TaskRepositoryGorm) GetTask(id int) (models.Task, error) {
	task, err := getByID(id, *r.db)
	if err != nil {
		return models.Task{}, err
	}
	task.Category = r.GetCategoryByID(int(task.CategoryID))
	fmt.Print(task.CategoryID)
	return task, nil
}

func (r *TaskRepositoryGorm) GetTasks(userID int) ([]models.Task, error) {
	var tasks []models.Task
	result := r.db.Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		logrus.Error(result.Error)
		return nil, result.Error
	}
	return tasks, nil
}
func getByID(id int, db gorm.DB) (models.Task, error) {
	var task models.Task
	result := db.Where("id = ?", id).First(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}

// TODO: remake
func (r *TaskRepositoryGorm) CreateCategory(title string) {
	result := r.db.Create(&models.Category{
		Title: title,
	})
	if result.Error != nil {
		return
	}
}

// TODO: remake
func (r *TaskRepositoryGorm) GetCategoryByID(id int) models.Category {
	var category models.Category
	result := r.db.Where("id = ?", id).First(&category)
	if result.Error != nil {
		return models.Category{}
	}
	return category
}

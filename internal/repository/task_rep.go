package repository

import (
	"errors"

	"github.com/DavidJackso/TodoApi/internal/lib/errs"
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TaskRepositoryGorm struct {
	db *gorm.DB
}

func NewTaskRepositoryGorm(dbgorm *gorm.DB) *TaskRepositoryGorm {
	return &TaskRepositoryGorm{
		db: dbgorm,
	}
}

func (r *TaskRepositoryGorm) CreateTask(task models.Task, userID uint) (uint, error) {
	task.UserID = userID

	

	result := r.db.Create(&task)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error add new task in gorm")
		return 0, result.Error
	}
	return task.ID, nil
}


func (r *TaskRepositoryGorm) DeleteTask(id uint, userID uint) error {
	task, err := getByID(id, r.db)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Error("task not found")
		return errs.ErrTaskNotFound
	}
	if task.UserID != userID {
		logrus.WithFields(logrus.Fields{
			"user_id":  userID,
			"owner_id": task.UserID,
		}).Warn("access denied")
		return errs.ErrAccessDenied
	}
	result := r.db.Delete(&task)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("failed delete task")
		return result.Error
	}
	return nil
}

func (r *TaskRepositoryGorm) UpdateTask(id uint, userID uint, updTask models.Task) (models.Task, error) {
	oldTask, err := getByID(id, r.db)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Error("task not found")
		return models.Task{}, errs.ErrTaskNotFound
	}

	if oldTask.UserID != userID {
		logrus.WithFields(logrus.Fields{
			"user_id":  userID,
			"owner_id": oldTask.UserID,
		}).Warn("access denied")
		return models.Task{}, errs.ErrAccessDenied
	}

	oldTask.Title = updTask.Title
	oldTask.Description = updTask.Description
	oldTask.CategoryID = updTask.CategoryID

	result := r.db.Save(&oldTask)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("failed update task")
		return models.Task{}, result.Error
	}

	return oldTask, nil
}

func (r *TaskRepositoryGorm) GetTask(id uint, userID uint) (models.Task, error) {
	var task models.Task
	result := r.db.Preload("Tag").Preload("Category").First(&task, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logrus.Error("task not found")
		return models.Task{}, errs.ErrTaskNotFound
	}

	if task.UserID != userID {
		logrus.WithFields(logrus.Fields{
			"user_id":  userID,
			"owner_id": task.UserID,
		}).Warn("access denied")
		return models.Task{}, errs.ErrAccessDenied
	}

	if result.Error != nil {
		logrus.WithError(result.Error).Error("failed get task")
		return models.Task{}, result.Error
	}

	return task, nil
}

func (r *TaskRepositoryGorm) GetTasks(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	result := r.db.Preload("Category").Preload("Tag").Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("failed get tasks")
		return nil, result.Error
	}
	return tasks, nil
}
func getByID(id uint, db *gorm.DB) (models.Task, error) {
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
func (r *TaskRepositoryGorm) GetCategoryByID(id uint) models.Category {
	var category models.Category
	result := r.db.Where("id = ?", id).First(&category)
	if result.Error != nil {
		return models.Category{}
	}
	return category
}

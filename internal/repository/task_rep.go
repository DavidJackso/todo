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

// TODO: переделать
func (r *TaskRepositoryGorm) CreateTask(task models.Task, userID uint) (uint, error) {
	task.UserID = userID

	categoryID, err := createCategory(task.CategoryID, task.Category, r.db)
	if err != nil {
		return 0, err
	}

	logrus.Error(categoryID)
	task.CategoryID = categoryID

	//TODO: вынести в отдельную функцию
	if len(task.Tags) != 0 {
		for i, tag := range task.Tags {
			var existingTag models.Tag
			if err := r.db.Where("title = ?", tag.Title).First(&existingTag).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					if err := r.db.Create(&tag).Error; err != nil {
						logrus.WithError(err).Error("error creating tag")
						return 0, err
					}
					task.Tags[i] = tag
				} else {
					logrus.WithError(err).Error("error checking tag existence")
					return 0, err
				}
			} else {
				task.Tags[i] = existingTag
			}
		}
	}

	result := r.db.Create(&task)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("error add new task")
		return 0, result.Error
	}
	return task.ID, nil
}

func createCategory(categoryID uint, categoryNew models.Category, db *gorm.DB) (uint, error) {
	var category models.Category
	result := db.First(&category, categoryID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if categoryNew.Title == "" {
			logrus.WithError(result.Error).Error("empty catalog")
			return 0, errs.ErrEmptyCategory
		}
		result := db.Create(&categoryNew)
		if result.Error != nil {
			logrus.WithError(result.Error).Error("failed create new category")
			return 0, result.Error
		}
		return categoryNew.ID, nil
	}
	if result.Error != nil {
		logrus.WithError(result.Error).Error("failed find category")
		return 0, result.Error
	}
	return categoryID, nil
}

func (r *TaskRepositoryGorm) DeleteTask(id uint, userID uint) error {
	task, err := getByID(id, r.db)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Info("task not found")
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
		logrus.Info("task not found")
		return models.Task{}, errs.ErrTaskNotFound
	}

	if oldTask.UserID != userID {
		logrus.WithFields(logrus.Fields{
			"user_id":  userID,
			"owner_id": oldTask.UserID,
		}).Warn("access denied")
		return models.Task{}, errs.ErrAccessDenied
	}

	updateFields := map[string]interface{}{}
	if updTask.Title != "" {
		updateFields["title"] = updTask.Title
	}
	if updTask.Description != "" {
		updateFields["description"] = updTask.Description
	}
	if updTask.CategoryID != 0 {
		updateFields["category_id"] = updTask.CategoryID
	}

	r.db.Model(&oldTask).Updates(updateFields)
	result := r.db.Save(&oldTask)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("failed update task")
		return models.Task{}, result.Error
	}

	return oldTask, nil
}

func (r *TaskRepositoryGorm) GetTask(id uint, userID uint) (models.Task, error) {
	var task models.Task
	result := r.db.Preload("Tags").Preload("Category").First(&task, id)

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

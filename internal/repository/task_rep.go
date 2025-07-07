package repository

import (
	"errors"

	"github.com/DavidJackso/TodoApi/internal/lib/errs"
	"github.com/DavidJackso/TodoApi/internal/models"
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

	categoryID, err := r.ensureCategory(task.CategoryID, task.Category)
	if err != nil {
		return 0, err
	}
	task.CategoryID = categoryID

	tags, err := r.ensureTags(task.Tags)
	if err != nil {
		return 0, err
	}
	task.Tags = tags

	if err := r.db.Create(&task).Error; err != nil {
		return 0, err
	}
	return task.ID, nil
}

func (r *TaskRepositoryGorm) ensureCategory(categoryID uint, category models.Category) (uint, error) {
	if categoryID != 0 {
		var existing models.Category
		if err := r.db.First(&existing, categoryID).Error; err == nil {
			return existing.ID, nil
		}
	}
	if err := r.db.Where("title = ?", category.Title).FirstOrCreate(&category).Error; err != nil {
		return 0, err
	}
	return category.ID, nil
}

func (r *TaskRepositoryGorm) ensureTags(tags []models.Tag) ([]models.Tag, error) {
	var result []models.Tag
	for _, tag := range tags {
		var existingTag models.Tag
		if err := r.db.Where("title = ?", tag.Title).First(&existingTag).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := r.db.Create(&tag).Error; err != nil {
					return nil, err
				}
				result = append(result, tag)
			} else {
				return nil, err
			}
		} else {
			result = append(result, existingTag)
		}
	}
	return result, nil
}

func (r *TaskRepositoryGorm) DeleteTask(id uint, userID uint) error {
	task, err := getByID(id, r.db)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrTaskNotFound
	}
	if task.UserID != userID {
		return errs.ErrAccessDenied
	}
	result := r.db.Delete(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TaskRepositoryGorm) UpdateTask(id uint, userID uint, updTask models.Task) (models.Task, error) {
	oldTask, err := getByID(id, r.db)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Task{}, errs.ErrTaskNotFound
	}

	if oldTask.UserID != userID {
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

	result := r.db.Model(&oldTask).Updates(updateFields)
	if result.Error != nil {
		return models.Task{}, result.Error
	}

	return oldTask, nil
}

func (r *TaskRepositoryGorm) GetTask(id uint, userID uint) (models.Task, error) {
	var task models.Task
	result := r.db.Preload("Tags").Preload("Category").First(&task, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Task{}, errs.ErrTaskNotFound
	}

	if task.UserID != userID {
		return models.Task{}, errs.ErrAccessDenied
	}

	if result.Error != nil {
		return models.Task{}, result.Error
	}

	return task, nil
}

func (r *TaskRepositoryGorm) GetTasks(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	result := r.db.Preload("Category").Preload("Tags").Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
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

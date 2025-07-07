package repository

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(models.User) (uint, error)
	DeleteUser(uint) error
	GetUserByEmailAndPassword(string, string) (models.User, error)
	GetUserByID(uint) (models.User, error)
	UpdateUser(uint, models.User) (models.User, error)
}

type TaskRepository interface {
	CreateTask(models.Task, uint) (uint, error)
	GetTask(uint, uint) (models.Task, error)
	DeleteTask(uint, uint) error
	UpdateTask(uint, uint, models.Task) (models.Task, error)
	GetTasks(uint) ([]models.Task, error)
}

type Repository struct {
	UserRepository
	TaskRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepositoryGorm(db),
		TaskRepository: NewTaskRepositoryGorm(db),
	}
}

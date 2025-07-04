package repository

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(models.User) (int, error)
	DeleteUser(int) error
	GetUser(string, string) (models.User, error)
	UpdateUser(models.User) error
}

type TaskRepository interface {
	CreateTask(models.Task, int) (int, error)
	GetTask(id int) (models.Task, error)
	DeleteTask(id int) error
	UpdateTask(id int) (models.Task, error)
	GetTasks() ([]models.Task, error)
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

package repository

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(models.User) (int, error)
	DeleteUser(int) error
	GetUser(string, string) (models.User, error)
	GetUserByID(int) (models.User, error)
	UpdateUser(int, models.User) (models.User, error)
}

type TaskRepository interface {
	CreateTask(models.Task, int) (int, error)
	GetTask(id int) (models.Task, error)
	DeleteTask(int) error
	UpdateTask(int, models.Task) (models.Task, error)
	GetTasks(int) ([]models.Task, error)
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

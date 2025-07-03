package service

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
)

type authorization interface {
	Regestration(models.User) (int, error)
	Authtorization(id int) (string, error)
}

type todoService interface {
	CreateTask(models.Task) (int, error)
	GetTask(int) (models.Task, error)
	DeleteTask(int) error
	UpdateTask(int) (models.Task, error)
	GetAllTask() ([]models.Task, error)
}

type Service struct {
	authorization
	todoService
}

func NewService(db *repository.Repository) *Service {
	return &Service{
		authorization: NewAuthorizationService(db),
		todoService:   NewTodoService(db),
	}
}

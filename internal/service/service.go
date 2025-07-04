package service

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
)

type authorization interface {
	CreateNewUser(models.User) (int, error)
	GenerateToken(string, string) (string, error)
	ParserToken(string) (int, error)
}
type todoService interface {
	CreateTask(models.Task) (int, error)
	GetTask(int) (models.Task, error)
	DeleteTask(int) error
	UpdateTask(int) (models.Task, error)
	GetTasks() ([]models.Task, error)
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

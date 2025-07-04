package service

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
)

type authorization interface {
	CreateNewUser(models.User) (int, error)
	GenerateToken(string, string) (string, error)
	ParseToken(string) (int, error)
}
type todoService interface {
	CreateTask(models.Task, int) (int, error)
	GetTask(int) (models.Task, error)
	DeleteTask(int) error
	UpdateTask(int, models.Task) (models.Task, error)
	GetTasks(id int) ([]models.Task, error)
}

type profileService interface {
	GetProfile(id int) (models.User, error)
	DeleteProfile(int) error
	UpdateProfile(int, models.User) (models.User, error)
}

type Service struct {
	authorization
	todoService
	profileService
}

func NewService(db *repository.Repository) *Service {
	return &Service{
		authorization:  NewAuthorizationService(db),
		todoService:    NewTodoService(db),
		profileService: NewProfileService(db),
	}
}

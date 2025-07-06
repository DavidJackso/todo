package service

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
)

type authorization interface {
	CreateNewUser(models.User) (uint, error)
	GenerateToken(string, string) (string, error)
	ParseToken(string) (uint, error)
}
type todoService interface {
	CreateTask(models.Task, uint) (uint, error)
	GetTask(uint, uint) (models.Task, error)
	DeleteTask(uint, uint) error
	UpdateTask(uint, uint, models.Task) (models.Task, error)
	GetTasks(uint) ([]models.Task, error)
}

type profileService interface {
	GetProfile(uint) (models.User, error)
	DeleteProfile(uint) error
	UpdateProfile(uint, models.User) (models.User, error)
}

type Services struct {
	authorization
	todoService
	profileService
}

func NewService(db *repository.Repository) *Services {
	return &Services{
		authorization:  NewAuthorizationService(db),
		todoService:    NewTodoService(db),
		profileService: NewProfileService(db),
	}
}

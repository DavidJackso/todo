package service

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
)

type IAuthorization interface {
	Regestration(models.User) (int, error)
	Authtorization(id int) (string, error)
}

type Service struct {
	IAuthorization
}

func NewService(db *repository.Repository) *Service {
	return &Service{
		IAuthorization: NewAuthorizationService(db),
	}
}

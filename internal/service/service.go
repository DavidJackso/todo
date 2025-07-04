package service

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
)

type Authorization interface {
	CreateNewUser(models.User) (int, error)
	GenerateToken(string, string) (string, error)
	ParserToken(string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(db *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(db),
	}
}

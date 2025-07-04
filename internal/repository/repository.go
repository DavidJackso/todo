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

type Repository struct {
	UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepositoryGorm(db),
	}
}

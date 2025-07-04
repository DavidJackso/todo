package repository

import (
	"errors"

	"github.com/DavidJackso/TodoApi/internal/lib/errs"
	"github.com/DavidJackso/TodoApi/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepositoryGorm(dbgorm *gorm.DB) *UserRepositoryGorm {
	return &UserRepositoryGorm{
		db: dbgorm,
	}
}

func (r *UserRepositoryGorm) CreateUser(user models.User) (int, error) {
	result := r.db.Create(&user)
	return int(user.ID), result.Error
}

func (r *UserRepositoryGorm) GetUser(email, password string) (models.User, error) {
	var user models.User

	result := r.db.Where("email=?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errs.ErrInvaliEmailorPassword
	}

	return user, result.Error
}

// TODO:Implement me
func (r *UserRepositoryGorm) DeleteUser(id int) error {
	return nil
}

// TODO: Implement me
func (r *UserRepositoryGorm) UpdateUser(models.User) error {
	return nil
}

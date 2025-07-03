package repository

import (
	"fmt"

	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/sirupsen/logrus"
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
	if result.Error != nil {
		logrus.Error("user not created", fmt.Sprintf("err:%s", result.Error))
		return 0, result.Error
	}
	return int(user.ID), nil
}

// TODO:Implement me
func (r *UserRepositoryGorm) GetUser(id int) (models.User, error) {
	return models.User{}, nil
}

// TODO:Implement me
func (r *UserRepositoryGorm) DeleteUser(id int) error {
	return nil
}

// TODO: Implement me
func (r *UserRepositoryGorm) UpdateUser(models.User) error {
	return nil
}

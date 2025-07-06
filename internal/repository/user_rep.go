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

func (r *UserRepositoryGorm) GetUserByID(id uint) (models.User, error) {
	user, err := getUserByID(id, r.db)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func getUserByID(id uint, db *gorm.DB) (models.User, error) {
	var user models.User

	result := db.Where("id = ?", id).Find(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (r *UserRepositoryGorm) CreateUser(user models.User) (uint, error) {
	result := r.db.Create(&user)
	if result.Error == gorm.ErrDuplicatedKey {
		return 0, errs.ErrEmailIsReady
	}
	return user.ID, result.Error
}

func (r *UserRepositoryGorm) GetUser(email, password string) (models.User, error) {
	var user models.User

	result := r.db.Where("email=?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errs.ErrInvaliEmailOrPassword
	}

	return user, result.Error
}

func (r *UserRepositoryGorm) DeleteUser(id uint) error {
	user, err := getUserByID(id, r.db)
	if err != nil {
		return err
	}
	result := r.db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepositoryGorm) UpdateUser(id uint, user models.User) (models.User, error) {
	oldUser, err := getUserByID(id, r.db)
	if err != nil {
		return models.User{}, err
	}
	oldUser.Email = user.Email
	oldUser.Password = user.Password
	oldUser.Name = user.Name
	result := r.db.Save(&oldUser)
	if result.Error != nil {
		return models.User{}, err
	}
	return oldUser, nil
}

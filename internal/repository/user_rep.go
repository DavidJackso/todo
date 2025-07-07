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

	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (r *UserRepositoryGorm) CreateUser(user models.User) (uint, error) {
	result := r.db.Create(&user)

	if result.Error != nil {
		if errs.IsDuplicateError(result.Error) {
			return 0, errs.ErrEmailIsAlreadyExists
		}
		return 0, result.Error
	}

	return user.ID, result.Error
}

func (r *UserRepositoryGorm) GetUserByEmailAndPassword(email, password string) (models.User, error) {
	var user models.User

	result := r.db.Where("email=? AND password", email, password).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errs.ErrInvalidEmailOrPassword
		}
		return user, result.Error
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

	updFields := map[string]interface{}{}
	if user.Email != "" {
		updFields["email"] = user.Email
	}
	if user.Name != "" {
		updFields["name"] = user.Name
	}
	if user.Password != "" {
		updFields["password"] = user.Password
	}

	result := r.db.Model(&oldUser).Updates(updFields)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	oldUser, _ = getUserByID(id, r.db)
	return oldUser, nil
}

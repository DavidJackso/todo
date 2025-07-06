package service

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
	"github.com/sirupsen/logrus"
)

type ProfileService struct {
	rep *repository.Repository
}

func NewProfileService(repository *repository.Repository) *ProfileService {
	return &ProfileService{
		rep: repository,
	}
}

// TODO: небезопасно
func (s *ProfileService) GetProfile(userID uint) (models.User, error) {
	user, err := s.rep.GetUserByID(userID)
	if err != nil {
		logrus.WithError(err).Error("failed to get profile")
		return models.User{}, err
	}
	return user, nil
}

// TODO: небезопасно
func (s *ProfileService) DeleteProfile(id uint) error {
	err := s.rep.DeleteUser(id)
	if err != nil {
		logrus.WithError(err).Error("failed to delete profile")
		return err
	}
	return nil
}

func (s *ProfileService) UpdateProfile(id uint, user models.User) (models.User, error) {
	newUser, err := s.rep.UpdateUser(id, user)
	if err != nil {
		logrus.WithError(err).Error("failed to update profile")
		return models.User{}, err
	}
	return newUser, nil
}

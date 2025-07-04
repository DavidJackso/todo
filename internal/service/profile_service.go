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

func (s *ProfileService) GetProfile(id int) (models.User, error) {
	user, err := s.rep.GetUserByID(id)
	if err != nil {
		logrus.Error(err)
		return models.User{}, err
	}
	return user, nil
}

func (s *ProfileService) DeleteProfile(id int) error {
	err := s.rep.DeleteUser(id)
	if err != nil {
		logrus.Error(err)
	}
	return nil
}

func (s *ProfileService) UpdateProfile(id int, user models.User) (models.User, error) {
	newUser, err := s.rep.UpdateUser(id, user)
	if err != nil {
		logrus.Error(err)
	}
	return newUser, nil
}

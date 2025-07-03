package service

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
)

// TODO: Вынести в env
const salt = "spf"

type AuthorizationService struct {
	rep *repository.Repository
}

func NewAuthorizationService(db *repository.Repository) *AuthorizationService {
	return &AuthorizationService{
		rep: db,
	}
}

func (s *AuthorizationService) Regestration(user models.User) (int, error) {
	user.Password = generateHash(user.Password)
	id, err := s.rep.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *AuthorizationService) Authtorization(id int) (string, error) {
	return "", nil
}

func generateHash(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	io.WriteString(h, salt)
	return hex.EncodeToString(h.Sum(nil))
}

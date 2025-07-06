package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

// TODO: Вынести в env
const salt = "spf"
const signingKey = "avbnggfdcc"

// TODO: Вынести в конфиг
const tokenTTL = 12 * time.Hour

type AuthorizationService struct {
	rep *repository.Repository
}

func NewAuthorizationService(db *repository.Repository) *AuthorizationService {
	return &AuthorizationService{
		rep: db,
	}
}

func (s *AuthorizationService) CreateNewUser(user models.User) (uint, error) {
	user.Password = generateHash(user.Password)
	id, err := s.rep.CreateUser(user)

	if err != nil {
		logrus.WithError(err).Error("failed create user")
		return 0, err
	}
	return id, nil
}

func (s *AuthorizationService) ParseToken(tokenString string) (uint, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	id, ok := claims["user_id"].(float64)
	if !ok {
		fmt.Print(id)
		return 0, errors.New("invalid token payload")
	}

	return uint(id), nil
}

func (s *AuthorizationService) GenerateToken(email, password string) (string, error) {
	user, err := s.rep.GetUser(email, generateHash(password))
	if err != nil {
		logrus.Error(err)
		return "", err
	}

	tokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(tokenTTL).Unix(),
		"iat":     time.Now().Unix(),
	}
	fmt.Print(tokenClaims)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims).SignedString([]byte(signingKey))
	if err != nil {
		logrus.Error(err)
	}
	return token, err
}

func generateHash(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	io.WriteString(h, salt)
	return hex.EncodeToString(h.Sum(nil))
}

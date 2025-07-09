package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type AuthorizationService struct {
	rep        *repository.Repository
	salt       string
	signingKey string
	tokenTTL   time.Duration
}

func NewAuthorizationService(db *repository.Repository) *AuthorizationService {
	ttl, err := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	if err != nil {
		ttl = 6
	}

	return &AuthorizationService{
		rep:        db,
		salt:       os.Getenv("SALT"),
		signingKey: os.Getenv("SIGNING_KEY"),
		tokenTTL:   time.Duration(ttl) * time.Hour,
	}
}

func (s *AuthorizationService) CreateNewUser(user models.User) (uint, error) {
	user.Password = generateHash(user.Password, s.salt)
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
		return []byte(s.signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	id, ok := claims["user_id"].(float64)
	if !ok {
		logrus.Error("failed get user id")
		return 0, errors.New("invalid token payload")
	}

	return uint(id), nil
}

func (s *AuthorizationService) GenerateToken(email, password string) (string, error) {
	user, err := s.rep.GetUserByEmailAndPassword(email, generateHash(password, s.salt))
	if err != nil {
		logrus.WithError(err).Error("failed get user")
		return "", err
	}

	tokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(s.tokenTTL).Unix(),
		"iat":     time.Now().Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims).SignedString([]byte(s.signingKey))
	if err != nil {
		logrus.WithError(err).Error("failed create jwt claims")
	}
	return token, err
}

func generateHash(password, salt string) string {
	h := md5.New()
	io.WriteString(h, password)
	io.WriteString(h, salt)
	return hex.EncodeToString(h.Sum(nil))
}

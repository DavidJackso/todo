package handler

import (
	"errors"
	"net/http"

	"github.com/DavidJackso/TodoApi/internal/lib/errs"
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
		return
	}

	id, err := h.service.CreateNewUser(user)
	if err != nil {
		if errors.Is(err, errs.ErrEmailIsAlreadyExists) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

type signInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed parse request body"})
		return
	}

	jwt, err := h.service.GenerateToken(input.Email, input.Password)
	if errors.Is(err, errs.ErrInvalidEmailOrPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed generate jwt token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

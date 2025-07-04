package handler

import (
	"errors"
	"net/http"

	"github.com/DavidJackso/TodoApi/internal/lib/errs"
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		logrus.Info("bad request", err)
		c.JSON(http.StatusBadRequest, map[string]int{
			"id": 0,
		})
		return
	}
	id, err := h.service.CreateNewUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]int{
			"id": 0,
		})
		return
	}

	logrus.Info("Success sign up")

	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	err := c.Bind(&input)
	if err != nil {
		logrus.Info("bad request", err)
		c.JSON(http.StatusBadRequest, "bad request")
	}

	jwt, err := h.service.GenerateToken(input.Email, input.Password)
	if errors.Is(err, errs.ErrInvaliEmailorPassword) {
		c.JSON(http.StatusBadRequest, "invalid email or password")
		return
	}

	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, "bad day")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": jwt,
	})
}

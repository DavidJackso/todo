package handler

import (
	"net/http"

	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		logrus.Info("bad request", err)
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	id, err := h.service.Regestration(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, id)
	}
}

func (h *Handler) signIn(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		logrus.Info("bad request", err)
		c.JSON(http.StatusBadRequest, "bad requet")
	}
}

package handler

import (
	"net/http"

	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateTask(c *gin.Context) {
	var task models.Task

	err := c.Bind(&task)

	if err != nil {
		logrus.Info("Bad request")
		c.JSON(http.StatusBadRequest, map[string]int{
			"id": 0,
		})
		return
	}

	id, err := h.service.CreateTask(task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]int{
			"id": 0,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})

}

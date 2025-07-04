package handler

import (
	"net/http"
	"strconv"

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

func (h *Handler) GetTask(c *gin.Context) {
	id := c.Param("id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		logrus.Info("bad request")
		c.JSON(http.StatusBadRequest, map[int]string{})
		return
	}

	task, err := h.service.GetTask(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[int]models.Task{})
		return
	}

	c.JSON(http.StatusOK, map[uint]models.Task{
		task.ID: task,
	})

}

func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		logrus.Info("Bad parametr")
		c.JSON(http.StatusBadRequest, "Badddddd")
		return
	}

	err = h.service.DeleteTask(uid)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, "bad")
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) GetAllTask(c *gin.Context) {
	return
}

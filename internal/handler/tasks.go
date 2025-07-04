package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		logrus.Error(err)
	}

	var task models.Task

	err = c.Bind(&task)

	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusBadRequest, map[string]int{
			"id": 0,
		})
		return
	}

	id, err := h.service.CreateTask(task, userID)

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

func (h *Handler) GetTasks(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		logrus.Error("bad")
	}

	tasks, err := h.service.GetTasks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "bad")
		return
	}
	fmt.Print(tasks)
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) UpdateTask(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var taskUpd models.Task

	c.Bind(taskUpd)

	task, err := h.service.UpdateTask(id, taskUpd)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "ok")
		return
	}

	fmt.Print(task)
	c.JSON(http.StatusOK, task)
}

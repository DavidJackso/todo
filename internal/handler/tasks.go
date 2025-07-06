package handler

import (
	"fmt"
	"net/http"

	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	var task models.Task

	err = c.Bind(&task)

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed parse request body"})
		return
	}

	id, err := h.service.CreateTask(task, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// TODO: добавить обработку ошибки задача не найдена
func (h *Handler) GetTask(c *gin.Context) {
	id, err := getTaskID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	task, err := h.service.GetTask(id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed get task"})
		return
	}

	c.JSON(http.StatusOK, task)

}

func (h *Handler) DeleteTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		logrus.WithError(err).Error("failed to get user ID")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	id, err := getTaskID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed get task ID"})
		return
	}

	err = h.service.DeleteTask(id, userID)
	if err != nil {
		logrus.WithError(err).Error("failed delete task")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task success deleted"})
}

func (h *Handler) GetTasks(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		logrus.WithError(err).Error("failed to get user ID")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	tasks, err := h.service.GetTasks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed get tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id, err := getTaskID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed parse task ID"})
		return
	}
	var taskUpd models.Task

	err = c.Bind(&taskUpd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed parse request body"})
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	task, err := h.service.UpdateTask(id, userID, taskUpd)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

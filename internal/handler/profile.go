package handler

import (
	"net/http"

	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteProfile(c *gin.Context) {
	id, err := getUserID(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "none useID")
		return
	}

	err = h.service.DeleteProfile(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "completed")
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	id, err := getUserID(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "none useID")
		return
	}

	var newUser models.User

	err = c.Bind(&newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, "failed decode body")
		return
	}

	updateUser, err := h.service.UpdateProfile(id, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed update user")
		return
	}

	c.JSON(http.StatusOK, updateUser)
}

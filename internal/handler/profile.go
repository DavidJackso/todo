package handler

import (
	"net/http"

	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteProfile(c *gin.Context) {
	id, err := getUserID(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	err = h.service.DeleteProfile(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete account"})
		return
	}

	c.JSON(http.StatusOK, "account success deleted")
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	id, err := getUserID(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	var newUser models.User

	err = c.Bind(&newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
		return
	}

	updateUser, err := h.service.UpdateProfile(id, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, updateUser)
}

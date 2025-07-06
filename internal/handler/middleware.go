package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const authorization = "Authorization"

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorization)
	if header == "" {
		c.JSON(http.StatusUnauthorized, "empty auth header")
		c.Abort()
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, "invalid auth header")
		c.Abort()
		return
	}

	userID, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusUnauthorized, "invalid jwt token")
		c.Abort()
		return
	}

	c.Set("userID", userID)
}

func getUserID(c *gin.Context) (uint, error) {
	id, ok := c.Get("userID")
	if !ok {
		return 0, errors.New("failde get userID")
	}

	return id.(uint), nil
}

func getTaskID(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Info("failed get task ID")
		return 0, err
	}
	return uint(id), nil
}

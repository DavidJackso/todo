package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const authorization = "Authorization"

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorization)
	if header == "" {
		c.JSON(http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, "invalid auth header")
		return
	}

	userID, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusUnauthorized, "invalid jwt token")
		return
	}

	c.Set("userID", userID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get("userID")
	if !ok {
		return 0, errors.New("failde get userID")
	}

	return id.(int), nil
}

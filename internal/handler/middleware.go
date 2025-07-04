package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

	userID, err := h.service.ParserToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, "invalid jwt token")
		return
	}

	c.Set("userID", userID)
}

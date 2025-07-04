package handler

import (
	"github.com/DavidJackso/TodoApi/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHanlder(s *service.Service) Handler {
	return Handler{
		service: s,
	}
}

func (h *Handler) InitRouting() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("sign-in", h.signIn)
		auth.POST("sign-up", h.signUp)
	}
	api := router.Group("/api", h.UserIdentity)
	{
		user := api.Group("/profile")
		{
			user.GET("/")
			user.PATCH("/")
			user.DELETE("/")
		}

		tasks := api.Group("/task")
		{
			tasks.POST("/")
			tasks.GET("/:id")
			tasks.GET("/")
			tasks.DELETE("/:id")
		}
	}
	return router
}

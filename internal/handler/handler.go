package handler

import (
	"github.com/DavidJackso/TodoApi/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Services
}

func NewHandler(s *service.Services) Handler {
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
			user.GET("/", h.GetProfile)
			user.PATCH("/", h.UpdateProfile)
			user.DELETE("/", h.DeleteProfile)
		}

		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.CreateTask)
			tasks.GET("/", h.GetTasks)
			tasks.GET("/:id", h.GetTask)
			tasks.PATCH("/:id", h.UpdateTask)
			tasks.DELETE("/:id", h.DeleteTask)
		}

	}
	return router
}

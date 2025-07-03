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
		auth.POST("sig-in")
		auth.POST("sign-up", h.signUp)
	}

	user := router.Group("/profile")
	{
		user.GET("/")
		user.PATCH("/")
		user.DELETE("/")
	}

	tasks := router.Group("/task")
	{
		tasks.POST("/", h.CreateTask)
		tasks.GET("/")
		tasks.GET("/:id", h.GetTask)
		tasks.PATCH("/")
		tasks.DELETE("/:id", h.DeleteTask)
	}
	return router
}

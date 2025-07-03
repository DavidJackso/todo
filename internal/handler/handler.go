package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHanlder() Handler {
	return Handler{}
}

func InitRouting() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("sig-in")
		auth.POST("sign-up")
	}

	user := router.Group("/profile")
	{
		user.GET("/")
		user.PATCH("/")
		user.DELETE("/")
	}

	tasks := router.Group("/task")
	{
		tasks.POST("/")
		tasks.GET("/:id")
		tasks.GET("/")
		tasks.DELETE("/:id")
	}
	return router
}

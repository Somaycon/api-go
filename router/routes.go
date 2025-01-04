package router

import (
	"github.com/Somaycon/api-go/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/")
	{
		// Task
		v1.GET("/task", handler.ShowTaskHandler)
		v1.POST("/task", handler.CreateTaskHandler)
		v1.DELETE("/task", handler.DeleteTaskHandler)
		v1.PUT("/task", handler.UpdateTaskHandler)
		v1.GET("/tasks", handler.ShowAllTasksHandler)

		// User
		v1.GET("/user")
		v1.GET("/users")
		v1.POST("user", handler.CreateUserHandler)
		v1.PUT("/user")
		v1.DELETE("/user")
	}
}

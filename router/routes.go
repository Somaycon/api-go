package router

import (
	"github.com/Somaycon/api-go/handler"
	"github.com/Somaycon/api-go/middleware"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/")
	{
		//Login and Register
		v1.POST("/login", handler.LoginUserHandler)
		v1.POST("/register", handler.CreateUserHandler)

		// Task
		v1.GET("/task", middleware.AuthMiddleware(), handler.ShowTaskHandler)
		v1.POST("/task", middleware.AuthMiddleware(), handler.CreateTaskHandler)
		v1.DELETE("/task", middleware.AuthMiddleware(), handler.DeleteTaskHandler)
		v1.PUT("/task", middleware.AuthMiddleware(), handler.UpdateTaskHandler)
		v1.GET("/tasks", middleware.AuthMiddleware(), handler.ShowAllTasksHandler)
		v1.GET("/user-tasks", middleware.AuthMiddleware(), handler.ShowUserTasksHandler)

		// User
		v1.GET("/users", middleware.AuthMiddleware(), handler.ShowAllUsers)
		v1.GET("/user", middleware.AuthMiddleware(), handler.ShowUserHandler)
		v1.PUT("/user", middleware.AuthMiddleware(), handler.UpdateUserHandler)
		v1.DELETE("/user", middleware.AuthMiddleware(), handler.DeleteUserHandler)

	}
}

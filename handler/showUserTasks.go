package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func ShowUserTasksHandler(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIdUint, ok := userId.(uint)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user id format"})
		return
	}
	tasks := []schemas.Task{}
	if err := db.Where("user_id = ?", userIdUint).Find(&tasks).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving tasks"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of tasks",
		"tasks":   tasks,
	})
}

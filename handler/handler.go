package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAllTasksHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET tasks",
	})
}
func ShowTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET tasks",
	})
}
func CreateTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST task",
	})
}
func DeleteTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE task",
	})
}
func UpdateTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "UPDATE task",
	})
}

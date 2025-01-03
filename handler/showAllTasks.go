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

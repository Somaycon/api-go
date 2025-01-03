package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "UPDATE task",
	})
}

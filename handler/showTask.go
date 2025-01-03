package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func ShowTaskHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusInternalServerError, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	task := schemas.Task{}
	if err := db.First(&task).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "task not found")
		return
	}
	SendSucess(ctx, "show-task", task)
}

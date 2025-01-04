package handler

import (
	"net/http"
	"strconv"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func ShowTaskHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusInternalServerError, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	taskId, err := strconv.Atoi(id)
	if err != nil {
		SendError(ctx, http.StatusBadRequest, "invalid format")
		return
	}
	task := schemas.Task{}
	if err := db.Where("id=?", taskId).First(&task).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "task not found")
		return
	}
	SendSucess(ctx, "show-task", task)
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteTaskHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	task := schemas.Task{}
	if err := db.First(&task, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("task with id: %s not found", id))
		return
	}
	if err := db.Delete(&task).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error to delete task with id: %s", id))
		return
	}
	SendSucess(ctx, "delete-task", task)
}
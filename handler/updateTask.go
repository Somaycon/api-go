package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateTaskHandler(ctx *gin.Context) {
	request := UpdateTaskRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errof("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	task := schemas.Task{}
	if err := db.First(&task, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "task not found")
		return
	}
	if request.Name != "" {
		task.Name = request.Name
	}
	if request.Description != "" {
		task.Description = request.Description
	}
	if err := db.Save(&task).Error; err != nil {
		logger.Errof("error to update task: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error to update task")
		return
	}
	SendSucess(ctx, "updated-task", task)
}

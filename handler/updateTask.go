package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/

// @Summary Update task
// @Descripition Update a task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param	id query string true "Task identification"
// @Success 200 {object} UpdateTaskResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /task [put]
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

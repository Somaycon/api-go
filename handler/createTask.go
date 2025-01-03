package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(ctx *gin.Context) {
	request := CreateTaskRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errof("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	task := schemas.Task{
		Name:        request.Name,
		Description: request.Description,
	}

	if err := db.Create(&task).Error; err != nil {
		logger.Errof("error creating task: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error creating task on database")
		return
	}
	SendSucess(ctx, "create-task", task)
}

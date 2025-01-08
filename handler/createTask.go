package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/

// @Summary Create task
// @Descripition Create a new task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param	request body CreateTaskRequest true "Request body"
// @Sucess 200 {object} CreateTaskResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /task [post]
func CreateTaskHandler(ctx *gin.Context) {
	request := CreateTaskRequest{}

	ctx.BindJSON(&request)
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
		return
	}
	userIdUint, ok := userId.(uint)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user Id format"})
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errof("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	task := schemas.Task{
		Name:        request.Name,
		Description: request.Description,
		UserId:      userIdUint,
	}

	if err := db.Create(&task).Error; err != nil {
		logger.Errof("error creating task: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error creating task on database")
		return
	}
	SendSucess(ctx, "create-task", task)
}

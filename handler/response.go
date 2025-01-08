package handler

import (
	"fmt"
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "aplication/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func SendSucess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "aplication/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s sucessfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateTaskResponse struct {
	Message string               `json:"message"`
	Data    schemas.TaskResponse `json:"data"`
}
type DeleteTaskResponse struct {
	Message string               `json:"message"`
	Data    schemas.TaskResponse `json:"data"`
}
type ShowTaskResponse struct {
	Message string               `json:"message"`
	Data    schemas.TaskResponse `json:"data"`
}
type ShowAllTasksResponse struct {
	Message string                 `json:"message"`
	Data    []schemas.TaskResponse `json:"data"`
}

package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func ShowAllTasksHandler(ctx *gin.Context) {
	tasks := []schemas.Task{}
	if err := db.Find(&tasks).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "error to get tasks")
		return
	}
	SendSucess(ctx, "List-tasks", tasks)

}

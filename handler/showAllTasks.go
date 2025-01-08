package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/

// @Summary List task
// @Descripition List all tasks
// @Tags Tasks
// @Accept json
// @Produce json
// @Success 200 {object} ShowAllTasksResponse
// @Failure 500 {object} ErrorResponse
// @Router /tasks [get]
func ShowAllTasksHandler(ctx *gin.Context) {
	tasks := []schemas.Task{}
	if err := db.Find(&tasks).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "error to get tasks")
		return
	}
	SendSucess(ctx, "List-tasks", tasks)

}

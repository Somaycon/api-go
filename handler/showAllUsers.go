package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/

// @Summary List users
// @Descripition List all users
// @Tags User
// @Accept json
// @Produce json
// @Param	token query string true "User token"
// @Success 200 {object} ShowAllUsersResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func ShowAllUsers(ctx *gin.Context) {
	users := []schemas.User{}
	if err := db.Find(&users).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "error to get users")
		return
	}
	SendSucess(ctx, "Show-all-Users", users)
}

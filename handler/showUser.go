package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/

// @Summary Get user
// @Descripition Get a user
// @Tags User
// @Accept json
// @Produce json
// @Param	token query string true "User token"
// @Success 200 {object} ShowUserResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /user [get]
func ShowUserHandler(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		SendError(ctx, http.StatusUnauthorized, "User not authenticated")
		return
	}

	user := schemas.User{}
	if err := db.Where("id=?", userId).First(&user).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "user not found")
		return
	}
	SendSucess(ctx, "Show-user", user)
}

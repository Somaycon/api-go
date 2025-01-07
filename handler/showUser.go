package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

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

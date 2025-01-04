package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func ShowAllUsers(ctx *gin.Context) {
	users := []schemas.User{}
	if err := db.Find(&users).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "error to get users")
		return
	}
	SendSucess(ctx, "Show-all-Users", users)
}

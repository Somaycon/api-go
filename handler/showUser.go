package handler

import (
	"net/http"
	"strconv"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func ShowUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusInternalServerError, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		SendError(ctx, http.StatusBadRequest, "invalid format")
		return
	}
	user := schemas.User{}
	if err := db.Where("id=?", userId).First(&user).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "user not found")
		return
	}
	SendSucess(ctx, "Show-user", user)
}

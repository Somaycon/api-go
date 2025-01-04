package handler

import (
	"fmt"
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	user := schemas.User{}
	if err := db.First(&user, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found", id))
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error to delete user with id: %s", id))
		return
	}
	SendSucess(ctx, "delete-user", user)
}

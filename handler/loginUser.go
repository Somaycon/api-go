package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginUserHandler(ctx *gin.Context) {
	request := LoginUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errof("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Find User
	user := schemas.User{}
	if err := db.Where("email=? AND password=?", request.Email, request.Password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Errof("user not found: %v", request.Email)
			SendError(ctx, http.StatusUnauthorized, "Invalid credentials")
		} else {
			logger.Errof("database error: %v", err.Error())
			SendError(ctx, http.StatusInternalServerError, "Error while fetching user")
		}
		return
	}
	SendSucess(ctx, "login-user", schemas.LoginUserResponse{
		Menssage: "Welcome back",
		UserId:   user.ID,
	})

}

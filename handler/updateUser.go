package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/

// @Summary Update user
// @Descripition Update a user
// @Tags User
// @Accept json
// @Produce json
// @Param	token query string true "User token"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [put]
func UpdateUserHandler(ctx *gin.Context) {
	request := UpdateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errof("validation error: %v", err)
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "int").Error())
		return
	}
	user := schemas.User{}
	if err := db.First(&user, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "user not found")
		return
	}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
	}
	if err := db.Save(&user).Error; err != nil {
		logger.Errof("error to update user: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error to update user")
		return
	}
	SendSucess(ctx, "updare-user", user)
}

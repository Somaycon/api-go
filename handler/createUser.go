package handler

import (
	"net/http"

	"github.com/Somaycon/api-go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/

// @Summary Create user
// @Descripition Create a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param	request body CreateUserRequest true "Request body"
// @Success 200 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errof("validation error: %v", err)
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	if err := db.Create(&user).Error; err != nil {
		logger.Errof("error to create user: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error to create user on database")
		return
	}
	SendSucess(ctx, "create-user", user)
}

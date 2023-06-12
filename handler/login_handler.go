package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/input"
	"github.com/katerji/UserAuthKit/service"
)

const LoginPath = "/auth/login"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	request := &LoginRequest{}
	err := c.BindJSON(request)
	if err != nil {
		sendBadRequest(c)
		return
	}
	authService := service.AuthService{}
	authInput := input.AuthInput{
		Email:    request.Email,
		Password: request.Password,
	}
	user, err := authService.Login(authInput)
	if err != nil {
		sendErrorMessage(c, err.Error())
		return
	}
	sendJSONResponse(c, user)
	return
}

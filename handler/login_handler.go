package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/input"
	"github.com/katerji/UserAuthKit/model"
	"github.com/katerji/UserAuthKit/service"
)

const LoginPath = "/auth/login"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  model.UserOutput `json:"user"`
	Token string           `json:"access_token"`
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
	token, err := service.CreateJwt(user)
	if err != nil {
		sendErrorMessage(c, "")
		return
	}
	response := LoginResponse{
		User:  user.ToOutput(),
		Token: token,
	}
	sendJSONResponse(c, response)
	return
}

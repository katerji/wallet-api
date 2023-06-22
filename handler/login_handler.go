package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/input"
	"github.com/katerji/UserAuthKit/model"
	"github.com/katerji/UserAuthKit/service"
)

const LoginPath = "/login"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User         model.UserOutput `json:"user"`
	Token        string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
}

func LoginHandler(c *gin.Context) {
	request := &LoginRequest{}
	err := c.BindJSON(request)
	if err != nil {
		fmt.Println(err)
		sendBadRequest(c)
		return
	}
	authInput := input.AuthInput{
		Email:    request.Email,
		Password: request.Password,
	}
	response, err := login(authInput)
	if err != nil {
		sendBadRequestWithMessage(c, err.Error())
		return
	}
	sendJSONResponse(c, response)
	return
}

func login(authInput input.AuthInput) (LoginResponse, error) {
	authService := service.AuthService{}
	user, err := authService.Login(authInput)
	if err != nil {
		return LoginResponse{}, err
	}
	jwtService := service.JWTService{}
	token, err := jwtService.CreateJwt(user)
	if err != nil {
		return LoginResponse{}, err
	}
	refreshToken, err := jwtService.CreateRefreshJwt(user)
	if err != nil {
		return LoginResponse{}, err
	}
	return LoginResponse{
		User:         user.ToOutput(),
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

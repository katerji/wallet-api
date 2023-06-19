package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/input"
	"github.com/katerji/UserAuthKit/service"
)

const RegisterPath = "/register"

const (
	errorMessageEmailAlreadyExists = "Email already exists."
	userRegisteredSuccessfully     = "User registered successfully."
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(c *gin.Context) {
	request := &RegisterRequest{}
	err := c.BindJSON(request)
	if err != nil {
		sendBadRequest(c)
		return
	}
	if request.Email == "" || request.Username == "" || request.Password == "" {
		sendBadRequest(c)
		return
	}
	registerUserInput := input.AuthInput{
		Email:    request.Email,
		Password: request.Password,
		Username: request.Username,
	}
	userService := service.AuthService{}
	_, err = userService.Register(registerUserInput)
	if err != nil {
		fmt.Println(err)
		sendErrorMessage(c, errorMessageEmailAlreadyExists)
		return
	}
	sendResponseMessage(c, userRegisteredSuccessfully)
	return
}

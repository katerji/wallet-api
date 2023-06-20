package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/service"
	"github.com/katerji/UserAuthKit/utils"
)

const RegistrationValidatorPath = "/register/validate"

type RegistrationRequest struct {
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
}

func RegistrationValidatorHandler(c *gin.Context) {
	var request RegistrationRequest
	if err := c.BindJSON(&request); err != nil {
		sendBadRequest(c)
		return
	}
	if !utils.IsEmailValid(request.Email) {
		sendBadRequestWithMessage(c, "Invalid email")
		return
	}
	authService := service.AuthService{}
	exists := authService.DoesEmailExist(request.Email)
	if exists {
		sendBadRequestWithMessage(c, "Email exists")
		return
	}
	exists = authService.DoesUsernameExist(request.Username)
	if exists {
		sendBadRequestWithMessage(c, "Username exists")
		return
	}
	sendResponseMessage(c, "Valid")
}

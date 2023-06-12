package controller

import "github.com/gin-gonic/gin"

const RegisterPath = "/register"

func RegisterController(c *gin.Context) {
	c.String(200, "RegisterController")
	return
}

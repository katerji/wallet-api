package main

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/controller"
)

func main() {
	router := gin.Default()

	router.GET(controller.LandingPath, controller.LandingController)

	router.POST(controller.RegisterPath, controller.RegisterController)

}

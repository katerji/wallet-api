package main

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/controller"
	"github.com/katerji/UserAuthKit/db"
)

func main() {

	initDB()

	router := gin.Default()

	router.GET(controller.LandingPath, controller.LandingController)

	router.POST(controller.RegisterPath, controller.RegisterController)

}

func initDB() {
	client := db.GetDatabaseInstance()
	err := client.Ping()
	if err != nil {
		panic(err)
	}
}

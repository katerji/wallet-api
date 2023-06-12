package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/katerji/UserAuthKit/db"
	"github.com/katerji/UserAuthKit/handler"
	"github.com/katerji/UserAuthKit/middleware"
)

func main() {
	initEnv()
	initDB()
	initWebServer()
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func initDB() {
	client := db.GetDbInstance()
	err := client.Ping()
	if err != nil {
		panic(err)
	}
}

func initWebServer() {
	router := gin.Default()
	api := router.Group("/api")

	api.GET(handler.LandingPath, handler.LandingController)

	auth := api.Group("/auth")
	auth.POST(handler.RegisterPath, handler.RegisterHandler)
	auth.POST(handler.LoginPath, handler.LoginHandler)

	api.Use(middleware.GetAuthMiddleware())

	api.GET(handler.UserInfoPath, handler.UserInfoHandler)

	err := router.Run(":85")
	if err != nil {
		panic(err)
	}
}

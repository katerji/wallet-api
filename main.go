package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/katerji/UserAuthKit/db"
	"github.com/katerji/UserAuthKit/handler"
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

	router.GET(handler.LandingPath, handler.LandingController)
	router.POST(handler.RegisterPath, handler.RegisterController)
}

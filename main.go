package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/crons"
	"github.com/katerji/UserAuthKit/db"
	"github.com/katerji/UserAuthKit/handler"
	"github.com/katerji/UserAuthKit/middleware"
	"github.com/robfig/cron"
)

func main() {
	initDB()
	startCron()
	initWebServer()
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
	auth.POST(handler.RegistrationValidatorPath, handler.RegistrationValidatorHandler)
	auth.POST(handler.LoginPath, handler.LoginHandler)
	auth.POST(handler.RefreshTokenPath, handler.RefreshTokenHandler)

	api.Use(middleware.GetAuthMiddleware())

	api.GET(handler.TokensPath, handler.TokensHandler)
	api.GET(handler.FavoritesPath, handler.FavoritesHandler)
	api.POST(handler.FavoriteCreatePath, handler.FavoriteCreateHandler)
	api.DELETE(handler.FavoriteDeletePath, handler.FavoriteDeleteHandler)

	err := router.Run(":85")
	if err != nil {
		panic(err)
	}
}

func startCron() {
	c := cron.New()
	err := c.AddFunc("*/5 * * * *", crons.FetchAndStoreTokensFromAPI)
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
}

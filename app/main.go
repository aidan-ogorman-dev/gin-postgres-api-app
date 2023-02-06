package main

import (
	"diary_app/controller"
	"diary_app/database"
	"diary_app/middleware"

	model "diary_app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func serveApplication() {
	router := gin.Default()
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entries", controller.GetAllEntries)

	router.Run("0.0.0.0:8080")
}

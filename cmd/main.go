package main

import (
	"fmt"
	"jwt-auth/config"
	"jwt-auth/internal/handlers"
	"jwt-auth/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()

	r.POST("/login", handlers.LoginHandler)
	r.GET("/protected", middleware.AuthMiddleware(), handlers.ProtectedHandler)

	fmt.Println("Сервер запущен на порту", config.ServerPort)
	r.Run(":" + config.ServerPort)
}

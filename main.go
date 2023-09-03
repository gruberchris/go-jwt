package main

import (
	"github.com/gin-gonic/gin"
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", middleware.RequireAuth, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"env":     gin.Mode(),
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	err := r.Run()

	if err != nil {
		log.Fatal("Error attempting to start gin server")
	}
}

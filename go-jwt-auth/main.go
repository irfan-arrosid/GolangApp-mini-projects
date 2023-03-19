package main

import (
	"GolangApp-mini-projects/go-jwt-auth/controllers"
	"GolangApp-mini-projects/go-jwt-auth/initializers"
	"GolangApp-mini-projects/go-jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DbConnect()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.UserSignUp)
	r.POST("/login", controllers.UserLogin)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}

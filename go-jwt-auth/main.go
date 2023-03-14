package main

import (
	"GolangApp-mini-projects/go-jwt-auth/controllers"
	"GolangApp-mini-projects/go-jwt-auth/initializers"

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

	r.Run()
}

// See in 27:00 on the video

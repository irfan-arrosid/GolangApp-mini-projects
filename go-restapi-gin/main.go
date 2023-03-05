package main

import (
	"GolangApp-mini-projects/go-restapi-gin/controllers"
	"GolangApp-mini-projects/go-restapi-gin/initializers"
	"GolangApp-mini-projects/go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	models.DbConnect()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.PostsCreate)

	r.Run()
}

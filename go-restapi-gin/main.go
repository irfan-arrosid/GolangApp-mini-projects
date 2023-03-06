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

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	// Go to 25:06 on the video

	r.Run()
}

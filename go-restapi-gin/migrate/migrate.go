package main

import (
	"GolangApp-mini-projects/go-restapi-gin/initializers"
	"GolangApp-mini-projects/go-restapi-gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	models.DbConnect()
}

func main() {
	models.DB.AutoMigrate(&models.Post{})
}

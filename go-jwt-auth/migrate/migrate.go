package main

import (
	"GolangApp-mini-projects/go-jwt-auth/initializers"
	"GolangApp-mini-projects/go-jwt-auth/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DbConnect()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}

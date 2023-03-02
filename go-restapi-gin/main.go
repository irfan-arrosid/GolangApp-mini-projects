package main

import (
	"GolangApp-mini-projects/go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.DB.Connection()

	r.Run()
}

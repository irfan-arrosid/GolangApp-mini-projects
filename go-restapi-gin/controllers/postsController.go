package controllers

import (
	"GolangApp-mini-projects/go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data from req.body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := models.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return data

	c.JSON(200, gin.H{
		"post": post,
	})
}

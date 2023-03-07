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

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	models.DB.Find(&posts)

	// Respond posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id from param
	id := c.Param("id")

	// Get the posts
	var post models.Post
	models.DB.First(&post, id)

	// Respond post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id from param
	id := c.Param("id")

	// Get the post of req.body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Get the post were updating
	var post models.Post
	models.DB.First(&post, id)

	// Updating
	models.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

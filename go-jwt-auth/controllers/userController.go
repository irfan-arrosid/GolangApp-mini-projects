package controllers

import (
	"GolangApp-mini-projects/go-jwt-auth/initializers"
	"GolangApp-mini-projects/go-jwt-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserSignUp(c *gin.Context) {
	// Get the email and password from req.body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body.",
		})

		return
	}

	// Hast the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password.",
		})

		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user.",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Create a new user is success.",
		"user":    user,
	})

}

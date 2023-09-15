package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/models"
)

func GetAllUsers(c *gin.Context) {

	// >> Get all users from database
	var users []models.User
	initializers.DB.Find(&users)

	c.IndentedJSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User with that ID doesn't exist",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	// >> Login user to his account if account exist
	var body struct {
		Email    string
		Password string
	}

	c.Bind(&body)

	var user models.User
	result := initializers.DB.First(&user, body)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User with that credentials doesn't exist.",
		})
	}
}

func Register(c *gin.Context) {
	var body struct {
		Name     string
		Surname  string
		Email    string
		Password string
		City     string
	}
	c.Bind(&body)

	// Check if user already exists in database
	var existingUser models.User
	result := initializers.DB.Where("email = ?", body.Email).First(&existingUser)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "User already exists.",
		})
		return
	}

	// Insert new user into database
	user := models.User{Name: body.Name, Surname: body.Surname, Email: body.Email, City: body.City, Password: body.Password}
	result = initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account successfully registered.",
	})
}

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
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// >> Parse JSON request body and bind it to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	// >> Find user in the database
	var user models.User
	result := initializers.DB.First(&user, "email = ?", body.Email)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User with that credentials doesn't exist.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Register(c *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Surname  string `json:"surname" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		City     string `json:"city" binding:"required"`
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

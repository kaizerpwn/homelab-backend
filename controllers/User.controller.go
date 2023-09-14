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

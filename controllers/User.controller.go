package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/models"
)

func GetAllUsers(c *gin.Context) {

	// >> Get all users from database
	var users []models.User
	initializers.DB.Find(&users)

	c.IndentedJSON(200, users)
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/models"
)

func GetAllRooms(c *gin.Context) {

	// >> Get all rooms from database
	var rooms []models.Room
	initializers.DB.Find(rooms)

	c.IndentedJSON(http.StatusOK, rooms)
}

func GetAllRoomsByID(c *gin.Context) {

	// >> Get all rooms from database
	var room models.Room
	result := initializers.DB.Find(room)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Room with that ID doesn't exist.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, room)
}

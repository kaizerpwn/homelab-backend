package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/models"
)

// @Summary Get all Rooms
// @Description Fetches a list of all rooms from the database.
// @Tags rooms
// @Accept json
// @Produce json
// @Success 200 {array} models.Room
// @Router /api/rooms [get]
func GetAllRooms(c *gin.Context) {

	// >> Get all rooms from database
	var rooms []models.Room
	initializers.DB.Find(rooms)

	c.IndentedJSON(http.StatusOK, rooms)
}

// @Summary Get a Room by ID
// @Description Retrieve a room's information by its unique ID.
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path int true "Room ID" Format(int64)
// @Success 200 {object} models.Room
// @Failure 404 {string} string "Room with that ID doesn't exist."
// @Router /api/rooms/{id} [get]
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

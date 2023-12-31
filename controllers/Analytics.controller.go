package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
)

// @Summary Get Number of All Devices
// @Description Retrieves the total number of devices in the database.
// @Tags analytics
// @Accept json
// @Produce json
// @Success 200 {integer} int64
// @Router /api/analytics/devices [get]
func GetNumberOfAllDevices(c *gin.Context) {
	// >> Get number of all devices
	var deviceCount int64
	initializers.DB.Count(&deviceCount)

	c.IndentedJSON(http.StatusOK, deviceCount)
}

// @Summary Get Number of All Rooms
// @Description Retrieves the total number of rooms in the database.
// @Tags analytics
// @Accept json
// @Produce json
// @Success 200 {integer} int64
// @Router /api/analytics/rooms [get]
func GetNumberOfAllRooms(c *gin.Context) {
	// >> Get number of all rooms
	var deviceCount int64
	initializers.DB.Count(&deviceCount)

	c.IndentedJSON(http.StatusOK, deviceCount)
}

// @Summary Get Number of Active Devices
// @Description Retrieves the total number of active devices.
// @Tags analytics
// @Accept json
// @Produce json
// @Success 200 {integer} int64
// @Router /api/analytics/activedevices [get]
func GetNumberOfActiveDevices(c *gin.Context) {
	// >> Get number of all inactive devices
	var deviceCount int64
	initializers.DB.Count(&deviceCount)

	c.IndentedJSON(http.StatusOK, deviceCount)
}

// @Summary Get Number of All Houses
// @Description Retrieves the total number of houses.
// @Tags analytics
// @Accept json
// @Produce json
// @Success 200 {integer} int64
// @Router /api/analytics/houses [get]
func GetNumberOfAllHouses(c *gin.Context) {
	var housesCount int64
	initializers.DB.Count(&housesCount)

	c.IndentedJSON(http.StatusOK, housesCount)
}

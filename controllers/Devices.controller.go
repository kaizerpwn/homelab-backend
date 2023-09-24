package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/models"
)

// @Summary Get all Devices
// @Description Fetches a list of all devices from the database.
// @Tags devices
// @Accept json
// @Produce json
// @Success 200 {array} models.Device
// @Router /devices [get]
func GetAllDevices(c *gin.Context) {
	// Get all devices from database
	var devices []models.Device
	initializers.DB.Find(&devices)
	c.IndentedJSON(http.StatusOK, devices)
}

// @Summary Get a Device by ID
// @Description Retrieve a device's information by its unique ID.
// @Tags devices
// @Accept json
// @Produce json
// @Param id path int true "Device ID" Format(int64)
// @Success 200 {object} models.Device
// @Failure 404 {string} string "Device with that ID doesn't exist."
// @Router /devices/{id} [get]
func GetDeviceById(c *gin.Context) {
	// >> Get specific device with id given

	id := c.Param("id")

	var device models.Device
	result := initializers.DB.First(&device, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Device with that ID doesn't exist.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, device)
}

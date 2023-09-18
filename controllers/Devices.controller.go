package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/models"
)

func GetAllDevices(c *gin.Context) {
	// Get all devices from database
	var devices []models.Device
	initializers.DB.Find(&devices)
	c.IndentedJSON(http.StatusOK, devices)
}

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

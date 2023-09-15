package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
)

func GetNumberOfAllDevices(c *gin.Context) {
	// >> Get number of all devices
	var deviceCount int64
	initializers.DB.Count(&deviceCount)

	c.IndentedJSON(http.StatusOK, deviceCount)
}

func GetNumberOfAllRooms(c *gin.Context) {
	// >> Get number of all rooms
	var deviceCount int64
	initializers.DB.Count(&deviceCount)

	c.IndentedJSON(http.StatusOK, deviceCount)
}

func GetNumberOfActiveDevices(c *gin.Context) {
	// >> Get number of all inactive devices
	var deviceCount int64
	initializers.DB.Count(&deviceCount)

	c.IndentedJSON(http.StatusOK, deviceCount)
}

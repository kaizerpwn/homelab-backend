package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/models"
)

// @Summary Get House by ID
// @Description Retrieve a house's information by its unique ID
// @Tags houses
// @Accept json
// @Produce json
// @Param id path int true "House ID" Format(int64)
// @Success 200 {object} models.House
// @Failure 404 {string} string "Device with that ID doesn't exist."
// @Router /api/houses/{id} [get]
func GetHouseById(c *gin.Context) {
	id := c.Param("id")
	var house models.House

	result := initializers.DB.First(&house, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "House with that ID doesn't exist.",
		})
	}
}

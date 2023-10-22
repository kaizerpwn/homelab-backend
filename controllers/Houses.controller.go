package controllers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/models"
	"github.com/kaizerpwn/homelab-backend/utils"
)

// @Summary Get House by ID
// @Description Retrieve a house's information by its unique ID
// @Tags houses
// @Accept json
// @Produce json
// @Param id path int true "House ID" Format(int64)
// @Success 200 {object} models.House
// @Failure 404 {string} string "House with that ID doesn't exist."
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

// @Summary Create House
// @Description Create new house with relevant data
// @Tags houses
// @Accept json
// @Produce json
// @Param id path int true "House ID" Format(int64)
// @Success 200 {string} string "Successfully added new house."
// @Failure 404 {string} string "Internal server error."
// @Router /api/houses/{id} [post]
func CreateHouse(c *gin.Context) {
	var body struct {
		Address      string                `json:"address" binding:"required"`
		City         string                `json:"city" binding:"required"`
		Country      string                `json:"country" binding:"required"`
		ZipCode      string                `json:"zipcode" binding:"required"`
		Floors       int8                  `json:"floors" binding:"required"`
		SquareMeters float32               `json:"squareMeters" binding:"required"`
		Latitude     float32               `json:"latitude" binding:"required"`
		Longitude    float32               `json:"longitude" binding:"required"`
		Image        *multipart.FileHeader `form:"image" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	var house models.House

	var houseQueryData struct {
		Longitude float32
		Latitude  float32
		Address   string
		Floors    int8
	}

	houseQueryData.Address = body.Address
	houseQueryData.Longitude = body.Longitude
	houseQueryData.Latitude = body.Latitude
	houseQueryData.Floors = body.Floors

	checkHouseExist := initializers.DB.First(&house, houseQueryData)

	if checkHouseExist != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "House with that data already exist.",
		})
		return
	}

	err := utils.UploadImage(c, body.Image)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	}

}

package main

import (
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.House{})
	initializers.DB.AutoMigrate(&models.Room{})
	initializers.DB.AutoMigrate(&models.Device{})
}

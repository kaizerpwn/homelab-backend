package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/controllers"
	"github.com/kaizerpwn/homelab-backend/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// >> Users routes
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserByID)

	// >> Devices routes
	r.GET("/devices", controllers.GetAllDevices)
	r.GET("/devices/:id", controllers.GetDeviceById)

	// >> Devices routes
	r.GET("/rooms", controllers.GetAllRooms)
	r.GET("/rooms/:id", controllers.GetAllRoomsByID)

	r.Run()
}

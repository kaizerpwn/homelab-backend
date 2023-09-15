package main

import (
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	// >> Users routes
	r.GET("/api/users/:id", controllers.GetUserByID)
	r.GET("/api/users", controllers.GetAllUsers)
	r.POST("/api/users/register", controllers.Register)
	r.POST("/api/users/login", controllers.Login)

	// >> Devices routes
	r.GET("/api/devices", controllers.GetAllDevices)
	r.GET("/api/devices/:id", controllers.GetDeviceById)

	// >> Devices routes
	r.GET("/api/rooms", controllers.GetAllRooms)
	r.GET("/api/rooms/:id", controllers.GetAllRoomsByID)

	// >> Analytics routes
	r.GET("/api/analytics/rooms", controllers.GetNumberOfAllRooms)
	r.GET("/api/analytics/devices", controllers.GetNumberOfAllDevices)
	r.GET("/api/analytics/activedevices", controllers.GetNumberOfActiveDevices)

	r.Run()
}

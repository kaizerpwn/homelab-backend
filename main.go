package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/controllers"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/utils"
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
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/api/users/register", controllers.Register)
	r.POST("/api/users/login", controllers.Login)

	r.Use(utils.VerifyToken)

	// >> Users routes
	r.GET("/api/users/:id", utils.VerifyToken, controllers.GetUserByID)
	r.GET("/api/users", utils.VerifyToken, controllers.GetAllUsers)

	// >> Devices routes
	r.GET("/api/devices", utils.VerifyToken, controllers.GetAllDevices)
	r.GET("/api/devices/:id", utils.VerifyToken, controllers.GetDeviceById)

	// >> Devices routes
	r.GET("/api/rooms", utils.VerifyToken, controllers.GetAllRooms)
	r.GET("/api/rooms/:id", utils.VerifyToken, controllers.GetAllRoomsByID)

	// >> Analytics routes
	r.GET("/api/analytics/rooms", utils.VerifyToken, controllers.GetNumberOfAllRooms)
	r.GET("/api/analytics/devices", utils.VerifyToken, controllers.GetNumberOfAllDevices)
	r.GET("/api/analytics/activedevices", utils.VerifyToken, controllers.GetNumberOfActiveDevices)

	r.Run()
}

package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/controllers"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/utils"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/kaizerpwn/homelab-backend/docs"
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

	v1 := r.Group("/api")

	users := v1.Group("/users")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)

		users.GET("/:id", utils.VerifyToken, controllers.GetUserByID)
		users.GET("/", utils.VerifyToken, controllers.GetAllUsers)
	}

	houses := v1.Group("/houses")
	{
		houses.GET("/:id", utils.VerifyToken, controllers.GetHouseById)
	}

	rooms := v1.Group("/rooms")
	{
		rooms.GET("/", utils.VerifyToken, controllers.GetAllRooms)
		rooms.GET("/:id", utils.VerifyToken, controllers.GetAllRoomsByID)
	}

	devices := v1.Group("/devices")
	{
		devices.GET("/", utils.VerifyToken, controllers.GetAllDevices)
		devices.GET("/:id", utils.VerifyToken, controllers.GetDeviceById)
	}

	analytics := v1.Group("/analytics")
	{
		analytics.GET("/houses", utils.VerifyToken, controllers.GetNumberOfAllHouses)
		analytics.GET("/rooms", utils.VerifyToken, controllers.GetNumberOfAllRooms)
		analytics.GET("/devices", utils.VerifyToken, controllers.GetNumberOfAllDevices)
		analytics.GET("/activedevices", utils.VerifyToken, controllers.GetNumberOfActiveDevices)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}

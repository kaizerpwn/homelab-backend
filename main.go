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

	// >> Users controllers
	r.GET("/users", controllers.GetAllUsers)

	r.Run()
}

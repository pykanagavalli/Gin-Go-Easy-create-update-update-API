package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamtejasmane/go-jwt-gin-gonic/routes"

	"kagz/configs"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoutes(router)

	router.Run("localhost:6000")
}

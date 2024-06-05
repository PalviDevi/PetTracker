package main

import (
	"PetTracker/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine
	var rg *gin.RouterGroup
	// New Default GIN Router
	router = gin.New()
	rg = router.Group("/api")
	// routes for api calling
	routes.PetsTrackerRoutes(rg)

	// Start Gin server
	router.Run(":8080") // Listening server by default

}

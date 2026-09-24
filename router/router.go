package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Initialize router
	router := gin.Default()

	// Initialize routers
	initializeRoutes(router)

	// Run server
	router.Run(":3000")
}

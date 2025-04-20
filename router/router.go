package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	router := gin.Default()

	initializeRoutes(router)

	// Run server
	router.Run(":8080") // Listening on port :8080
}

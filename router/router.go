package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	// Initialize routes
	initializeRoutes(r)

	// Run the server on 0.0.0.0:8080
	r.Run("0.0.0.0:8080")
}

package main

import (
	"travel-destinations-api/config"
)

func main() {
	router := config.SetupServer()

	// Initialize routes
	routes.SetupDestinationRoutes(router)
	routes.SetupReviewRoutes(router)

	// Start the server
	router.Run("localhost:8080")
}

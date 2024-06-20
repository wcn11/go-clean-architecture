package main

import (
	"fmt"
	"go-clean-architecture/pkg/config"
	"go-clean-architecture/pkg/routes"
	"net/http"
)

func main() {
	// Initialize the configuration.
	config.ViperConfig()

	// Initialize the router
	r := routes.InitRouter()

	// Retrieve the server port from configuration
	port := config.ViperConfig().GetInt("server.port")
	fmt.Printf("Starting server on port %d...\n", port)

	// Start the HTTP server.
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
}

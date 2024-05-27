package main

import (
	"fmt"
	"full/pkg/config"
	"full/pkg/logger"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// main function initializes configuration, sets up the DI container, and starts the server.
func main() {
	// Initialize the configuration.
	config.InitConfig()

	// Build the DI container.
	container := config.BuildContainer()

	// Invoke the server startup with the necessary dependencies.
	err := container.Invoke(func(db *gorm.DB, router *mux.Router) {
		defer func() {
			sqlDb, err := db.DB()
			if err != nil {
				log.Fatal("Cannot close DB: ", err)
			}
			if err := sqlDb.Close(); err != nil {
				log.Fatal("Cannot close DB: ", err)
			}
		}()

		// Start the HTTP server.
		logger.Infof("Starting server on %d", config.AppConfig.Server.Port)
		logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.Server.Port), router))
	})

	// Handle any errors that occur during the Invoke process.
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

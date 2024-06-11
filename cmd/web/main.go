package main

import (
	"fmt"
	"github.com/gorilla/mux"
	clientGrpc "go-clean-architecture/internal/client/grpc"
	grpcServer "go-clean-architecture/internal/server/grpc"
	"go-clean-architecture/pkg/config"
	"go-clean-architecture/pkg/logger"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func main() {
	// Initialize the configuration.
	config.InitConfig()

	// Build the DI container.
	container := config.BuildContainer()

	// Start gRPC server in a goroutine.
	go func() {
		log.Println("Starting gRPC server...")
		grpcServer.StartGRPCServer()
	}()

	// Wait a bit to ensure the server is up.
	time.Sleep(2 * time.Second)

	// Call the gRPC client.
	clientGrpc.CallSayHello("world")

	// Invoke the server startup with the necessary dependencies.
	err := container.Invoke(func(db *gorm.DB, router *mux.Router) {
		// Ensure the DB connection is closed when the server shuts down.
		defer func() {
			sqlDb, err := db.DB()
			if err != nil {
				log.Fatal("Cannot get DB instance: ", err)
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

package config

import (
	"github.com/spf13/viper"
	"log"
)

func ViperConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./")

	// Set default values
	config.SetDefault("Server.Port", 8080)
	config.SetDefault("Database.Driver", "postgres")

	err := config.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found, using defaults and environment variables: %v", err)
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}

	// Optional: Log the effective configuration for debugging purposes
	log.Printf("Configuration loaded")
	return config
}

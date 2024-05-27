package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port     int
	Timezone string
}

type DatabaseConfig struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	Name     string
	Ssl      string
}

var AppConfig Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	// Adding multiple config paths
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")  // Add additional paths as needed
	viper.AddConfigPath("/etc/myapp") // Example for Unix-based systems

	// Set default values
	viper.SetDefault("Server.Port", 8080)
	viper.SetDefault("Database.Driver", "postgres")

	// Automatically override config with environment variables if set
	viper.AutomaticEnv()
	viper.SetEnvPrefix("myapp")
	viper.BindEnv("Server.Port", "MYAPP_SERVER_PORT")
	viper.BindEnv("Database.User", "MYAPP_DB_USER")
	viper.BindEnv("Database.Password", "MYAPP_DB_PASSWORD")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Config file not found, using defaults and environment variables: %v", err)
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	// Optional: Log the effective configuration for debugging purposes
	log.Printf("Configuration loaded: %+v", AppConfig)
}

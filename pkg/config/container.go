package config

import (
	"github.com/sirupsen/logrus"
	"go-clean-architecture/pkg/controllers"
	"go-clean-architecture/pkg/logger"
	userRepo "go-clean-architecture/pkg/repository/user"
	"go-clean-architecture/pkg/routes"
	emailService "go-clean-architecture/pkg/usecase/email"
	userService "go-clean-architecture/pkg/usecase/user"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Register the database provider
	err := container.Provide(GetDB)
	if err != nil {
		logger.Fatalf("failed to build container: %v", err)
	}

	// Automatically register providers
	registerProviders(container)

	return container
}

func registerProviders(container *dig.Container) {

	// List of provider functions, register here for dependency injection
	providers := []interface{}{
		func(db *gorm.DB) userRepo.IUserRepository {
			return userRepo.NewUserRepository(db)
		},
		func(repo userRepo.IUserRepository, emailSvc emailService.IEmailService) userService.IUserService {
			return userService.NewUserService(repo, emailSvc)
		},
		emailService.NewEmailService,
		controllers.NewIndexController,
		routes.InitRouter,
		ProvideLogger,
	}

	for _, provider := range providers {
		if err := container.Provide(provider); err != nil {
			logrus.Fatalf("failed to provide: %v", err)
		}
	}
}

// ProvideLogger register the logger
func ProvideLogger() interface{} {
	return func() *logrus.Logger {
		return logger.Logger
	}
}

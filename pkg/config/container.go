package config

import (
	"full/pkg/controllers"
	"full/pkg/logger"
	userRepo "full/pkg/repository/user"
	"full/pkg/routes"
	emailService "full/pkg/usecase/email"
	userService "full/pkg/usecase/user"
	"github.com/sirupsen/logrus"
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

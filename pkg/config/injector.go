//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"go-clean-architecture/pkg/controllers"
	UserRepository "go-clean-architecture/pkg/repository/user"
	"go-clean-architecture/pkg/uow"
	UserService "go-clean-architecture/pkg/usecase/user"
)

var group = wire.NewSet(
	ViperConfig,
	GetDB,
	wire.Bind(new(UserRepository.IUserRepository), new(*UserRepository.UserRepositoryImpl)),

	uow.NewUowImpl,
	wire.Bind(new(uow.UnitOfWork), new(*uow.UnitOfWorkImpl)),
	UserRepository.NewUserRepository,
	wire.Bind(new(UserService.IUserService), new(*UserService.UserServiceImpl)),
	UserService.NewUserService,
	controllers.NewUserController,
)

func InitInjector() *controllers.UserController {
	// Build the dependency graph using the wire set
	wire.Build(group)

	// Return nil because UserController instance is created and injected
	return nil
}

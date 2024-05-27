package user

import (
	"context"
	"full/pkg/domain"
	"full/pkg/dto"
)

type IUserService interface {
	GetUserById(ctx context.Context, id int) (*dto.UserResponse, error)
	CreateUser(ctx context.Context, userDto *dto.CreateUserRequest) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id int) error
}

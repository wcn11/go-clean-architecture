package user

import (
	"context"
	"full/pkg/domain"
)

type IUserRepository interface {
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}

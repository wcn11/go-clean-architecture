package user

import (
	"context"
	"go-clean-architecture/pkg/domain"
	"go-clean-architecture/pkg/dto"
	"go-clean-architecture/pkg/repository/user"
	"time"
)

type UserServiceImpl struct {
	UserRepository user.IUserRepository
}

func NewUserService(userRepo user.IUserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepo,
	}
}

func (u *UserServiceImpl) GetUserById(ctx context.Context, id int) (*dto.UserResponse, error) {
	return &dto.UserResponse{
		ID:        uint(id),
		Name:      "wcn",
		Phone:     "089123123",
		Gender:    1,
		Email:     "wewer@@gmai.cw",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, userDto *dto.CreateUserRequest) (*dto.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) UpdateUser(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

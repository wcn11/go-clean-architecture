package user

import (
	"context"
	"full/pkg/domain"
	"full/pkg/dto"
	"full/pkg/logger"
	"full/pkg/repository/user"
	"full/pkg/usecase/email"
	"time"
)

type UserServiceImpl struct {
	userRepository user.IUserRepository
	emailService   email.IEmailService
}

func NewUserService(repo user.IUserRepository, emailSvc email.IEmailService) IUserService {
	return &UserServiceImpl{
		userRepository: repo,
		emailService:   emailSvc,
	}
}

func (u *UserServiceImpl) GetUserById(ctx context.Context, id int) (*dto.UserResponse, error) {
	userData, err := u.userRepository.GetUserByID(ctx, id)

	if err != nil {
		logger.Fatalf("Error while fetching user by ID: %s", id)
	}

	return &dto.UserResponse{
		ID:        userData.ID,
		Name:      userData.Name,
		Phone:     userData.Phone,
		Gender:    userData.Gender,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}, nil
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, userDto *dto.CreateUserRequest) (*dto.UserResponse, error) {
	createUser := &domain.User{
		Name:      userDto.Name,
		Phone:     userDto.Phone,
		Gender:    userDto.Gender,
		Email:     userDto.Email,
		Password:  userDto.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userResponse, err := u.userRepository.CreateUser(ctx, createUser)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        userResponse.ID,
		Name:      userResponse.Name,
		Phone:     userResponse.Phone,
		Gender:    userResponse.Gender,
		Email:     userResponse.Email,
		CreatedAt: userResponse.CreatedAt,
		UpdatedAt: userResponse.UpdatedAt,
	}, nil
}

func (u *UserServiceImpl) UpdateUser(ctx context.Context, user *domain.User) error {
	panic("implement me")
}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, id int) error {
	// TODO: implement me
	panic("implement me")
}

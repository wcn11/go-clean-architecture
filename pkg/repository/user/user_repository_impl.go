package user

import (
	"context"
	"full/pkg/domain"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RepositoryImpl) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
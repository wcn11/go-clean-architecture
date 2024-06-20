package user

import (
	"context"
	"go-clean-architecture/pkg/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	//err := r.db.WithContext(ctx).First(&user, id).Error
	//if err != nil {
	//	return nil, err
	//}
	return &user, nil
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	//err := r.db.WithContext(ctx).Create(user).Error
	//if err != nil {
	//	return nil, err
	//}
	return user, nil
}

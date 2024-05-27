package dto

import "time"

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Gender   int16  `json:"gender" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Gender    int16     `json:"gender"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

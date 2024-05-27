package domain

import "time"

type User struct {
	ID         uint      `gorm:"primary_key;autoIncrement" json:"id"`
	Name       string    `gorm:"name"`
	Phone      string    `gorm:"phone"`
	Gender     int16     `gorm:"gender"`
	Email      string    `gorm:"email"`
	Password   string    `gorm:"password"`
	CreatedAt  time.Time `gorm:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at"`
	ModifiedAt time.Time `gorm:"modified_at"`
	DeletedAt  time.Time `gorm:"deleted_at"`
}

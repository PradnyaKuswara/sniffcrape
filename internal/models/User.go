package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id"`
	Username  string         `json:"username" gorm:"unique"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	AvatarURL *string        `json:"avatar_url"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}

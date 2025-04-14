package models

import "gorm.io/gorm"

type User struct {
	ID		uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

type UserMigration struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (UserMigration) TableName() string {
	return "users"
}


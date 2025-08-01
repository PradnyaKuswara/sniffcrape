package models

import "github.com/golang-jwt/jwt/v5"

type AuthRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthRegisterRequest struct {
	FirstName string  `json:"first_name" validate:"required"`
	LastName  string  `json:"last_name" validate:"required"`
	Username  string  `json:"username" validate:"required"`
	AvatarURL *string `json:"avatar_url"`
	Email     string  `json:"email" validate:"required,email"`
	Password  string  `json:"password" validate:"required,min=6"`
}

type JwtAttributes struct {
	ID        uint    `json:"id"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	AvatarURL *string `json:"avatar_url"`
	jwt.RegisteredClaims
}

type AuthUser struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Username  string  `json:"username"`
	AvatarURL *string `json:"avatar_url"`
	Email     string  `json:"email"`
}

type AuthUserWithID struct {
	UserID uint `json:"user_id"`
	AuthUser
}

type AuthResponse struct {
	Token string   `json:"token"`
	User  AuthUser `json:"user"`
}

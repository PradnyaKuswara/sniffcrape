package services

import (
	"time"

	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	customerrors "github.com/PradnyaKuswara/sniffcrape/pkg/errors"
	"github.com/PradnyaKuswara/sniffcrape/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserService *UserService
}

func NewAuthService(
	userService *UserService,
) *AuthService {
	return &AuthService{
		UserService: userService,
	}
}

func (s *AuthService) Login(email, password string) (*models.AuthResponse, error) {
	user, err := s.UserService.GetUserByEmail(email)
	if err != nil {
		err = customerrors.ErrInvalidCredentials
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, customerrors.ErrInvalidCredentials
	}

	token, err := utils.GenerateJWT(models.JwtAttributes{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		AvatarURL: user.AvatarURL,
	})

	if err != nil {
		return nil, err
	}

	response := &models.AuthResponse{
		Token: token,
		User: models.AuthUser{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			AvatarURL: user.AvatarURL,
			Email:     user.Email,
		},
	}
	return response, err
}

func (s *AuthService) Register(authReq models.AuthRegisterRequest) error {
	_, err := s.UserService.GetUserByEmail(authReq.Email)
	if err == nil {
		return customerrors.ErrConflict
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		FirstName: authReq.FirstName,
		LastName:  authReq.LastName,
		Username:  authReq.Username,
		AvatarURL: authReq.AvatarURL,
		Email:     authReq.Email,
		Password:  string(hashedPassword),
	}
	err = s.UserService.UserInterface.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) ValidateUser(tokenStr string) (*models.AuthUserWithID, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.JwtAttributes{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return utils.GetJWTSecret(), nil
	})

	if err != nil || token == nil {
		return nil, customerrors.ErrUnauthenticated
	}

	claims, ok := token.Claims.(*models.JwtAttributes)
	if !ok || claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, customerrors.ErrTokenExpired
	}

	return &models.AuthUserWithID{
		UserID: claims.ID,
		AuthUser: models.AuthUser{
			FirstName: claims.FirstName,
			LastName:  claims.LastName,
			Username:  claims.Username,
			AvatarURL: claims.AvatarURL,
			Email:     claims.Email,
		},
	}, nil
}

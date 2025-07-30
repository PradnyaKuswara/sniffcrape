package services

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	customerrors "github.com/PradnyaKuswara/sniffcrape/pkg/errors"
	"github.com/PradnyaKuswara/sniffcrape/pkg/utils"
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

func (s *AuthService) Login(email, password string) (*models.User, string, error) {
	user, err := s.UserService.GetUserByEmail(email)
	if err != nil {
		err = customerrors.ErrInvalidCredentials
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, "", customerrors.ErrInvalidCredentials
	}

	token, err := utils.GenerateJWT(user.ID)
	return user, token, nil
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

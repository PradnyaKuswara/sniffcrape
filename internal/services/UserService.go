package services

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"github.com/PradnyaKuswara/sniffcrape/internal/repositories"
)

type UserService struct {
	UserInterface repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) *UserService {
	return &UserService{UserInterface: repo}
}

func (s *UserService) GetUserByID(id string) (models.User, error) {
	user, err := s.UserInterface.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
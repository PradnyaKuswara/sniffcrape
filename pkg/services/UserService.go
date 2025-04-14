package services

import (
	"github.com/PradnyaKuswara/sniffcrape/pkg/models"
	"github.com/PradnyaKuswara/sniffcrape/pkg/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func (s *UserService) GetUserByID(id string) (models.User, error) {
	user, err := s.UserRepository.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
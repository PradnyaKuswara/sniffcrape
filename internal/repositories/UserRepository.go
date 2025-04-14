package repositories

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	GetUserByID(id string) (models.User, error)
	CreateUser(user models.User) error
}

func (r *UserRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(user models.User) error {
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

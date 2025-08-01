package repositories

import (
	"errors"

	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	customerrors "github.com/PradnyaKuswara/sniffcrape/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(user models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrDataNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user models.User) error {
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrDataNotFound
		}
		return nil, err
	}
	return &user, nil
}

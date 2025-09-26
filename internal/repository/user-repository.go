package repository

import (
	"context"

	"github.com/naol86/go/authGo/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// CreateUser implements domain.UserRepository.
func (u *UserRepository) CreateUser(c context.Context, user domain.UserSignupRequest) (domain.User, error) {
	newUser := domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := u.db.WithContext(c).Create(&newUser).Error; err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}

// GetUserByEmail implements domain.UserRepository.
func (u *UserRepository) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	err := u.db.WithContext(c).Where("email = ?", email).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// GetUserByID implements domain.UserRepository.
func (u *UserRepository) GetUserByID(c context.Context, id string) (domain.User, error) {
	var user domain.User
	err := u.db.WithContext(c).Where("id = ?", id).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

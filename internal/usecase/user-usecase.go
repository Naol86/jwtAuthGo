package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/naol86/go/authGo/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo domain.UserRepository
	Timeout  time.Duration
}

// Signin implements domain.UserUseCase.
func (u *UserUseCase) Signin(c context.Context, user domain.UserSigninRequest) (domain.User, error) {
	userData, err := u.userRepo.GetUserByEmail(c, user.Email)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, errors.New("invalid credentials")
	}
	return userData, nil
}

// Signup implements domain.UserUseCase.
func (u *UserUseCase) Signup(c context.Context, user domain.UserSignupRequest) (domain.User, error) {
	_, err := u.userRepo.GetUserByEmail(c, user.Email)
	if err == nil {
		return domain.User{}, errors.New("user already exists")
	}
	return u.userRepo.CreateUser(c, user)
}

func NewUserUseCase(userRepo domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
		Timeout:  timeout,
	}
}

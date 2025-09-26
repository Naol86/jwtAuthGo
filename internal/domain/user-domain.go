package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"size:255; not null" json:"name"`
	Email    string `gorm:"size:255; not null; unique" json:"email"`
	Password string `gorm:"size:255; not null" json:"-"`
}

type UserSignupRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserSigninRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserResponse struct {
	Message      string `json:"message"`
	Success      bool   `json:"success"`
	Data         User   `json:"data"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type UserRepository interface {
	CreateUser(c context.Context, user UserSignupRequest) (User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByID(c context.Context, id string) (User, error)
}

type UserUseCase interface {
	Signup(c context.Context, user UserSignupRequest) (User, error)
	Signin(c context.Context, user UserSigninRequest) (User, error)
}

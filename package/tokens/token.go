package tokens

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/naol86/go/authGo/internal/domain"
)

func CreateAccessToken(user *domain.User, secret string, exp int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(exp) * time.Hour)
	claims := domain.Claims{
		Email: user.Email,
		ID:    string(rune(user.ID)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func CreateRefreshToken(user *domain.User, secret string, exp int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(exp) * time.Hour)
	claims := domain.RefreshClaims{
		ID: string(rune(user.ID)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(token string, secret string) (bool, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}

	if !tokenString.Valid {
		return false, errors.New("invalid token")
	}
	return true, nil
}

func GetEmail(token string, secret string) (string, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", errors.New("invalid token")
	}

	claims, ok := tokenString.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("unexpected claims")
	}
	return claims["email"].(string), nil
}

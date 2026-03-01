package token

import (
	"errors"
	"os"
	"sentinel-apm-api/internal/modules/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	GenerateAccessToken(user *user.UserRequest) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type tokenManager struct {
	jwtSecret string
}

func NewTokenManager() *tokenManager {
	return &tokenManager{
		jwtSecret: os.Getenv("JWT_SECRET"),
	}
}

type TokenCliams struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.RegisteredClaims
}

func (tm *tokenManager) GenerateAccessToken(user *user.UserRequest) (string, error) {
	claims := &TokenCliams{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tm.jwtSecret))
}

func (tm *tokenManager) ValidateToken(tokenString string) (*TokenCliams, error) {
	claims := &TokenCliams{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(tm.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

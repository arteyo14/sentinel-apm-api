package auth

import (
	"errors"
	"sentinel-apm-api/internal/modules/user"
	"sentinel-apm-api/utils/context"
	"sentinel-apm-api/utils/token"
	"sentinel-apm-api/utils/validation"
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(request LoginRequest) (*LoginResponse, error)
	Register(request RegisterRequest) (*RegisterResponse, error)
	RefreshToken(request RefreshTokenRequest, sessionUser *context.SessionUser) (*RefreshTokenResponse, error)
	Logout(sessionUser *context.SessionUser) error
}

type authService struct {
	authRepo AuthRepository
}

func NewAuthService(authRepo AuthRepository) AuthService {
	return &authService{authRepo: authRepo}
}

func (s authService) Register(request RegisterRequest) (*RegisterResponse, error) {
	userData, err := s.authRepo.CheckUserEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if userData != nil {
		return nil, errors.New("user already exists")
	}

	var userRequest = user.User{
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	userId, err := s.authRepo.Register(&userRequest)
	if err != nil {
		return nil, err
	}

	return &RegisterResponse{
		UserID: userId.String(),
	}, nil
}

func (s authService) Login(request LoginRequest) (*LoginResponse, error) {
	userData, err := s.authRepo.CheckUserEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if err := validation.ValidatePassword(request.Password, userData.Password); err != nil {
		return nil, err
	}

	var userRequest = user.User{
		ID:        userData.ID,
		Email:     userData.Email,
		Password:  userData.Password,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
	}

	accessToken, err := token.NewTokenManager().GenerateAccessToken(&userRequest)
	if err != nil {
		return nil, err
	}

	refreshToken, err := token.NewTokenManager().GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	tokenData, err := s.authRepo.CreateRefreshToken(accessToken, refreshToken, userData.ID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  tokenData.AccessToken,
		RefreshToken: tokenData.RefreshToken,
		ExpiresAt:    tokenData.ExpiredAt,
	}, nil
}

func (s authService) RefreshToken(request RefreshTokenRequest, sessionUser *context.SessionUser) (*RefreshTokenResponse, error) {
	if request.ExpiredAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	parsedID, err := uuid.Parse(sessionUser.UserID)
	if err != nil {
		return nil, errors.New("invalid user id format")
	}

	accessToken, err := token.NewTokenManager().GenerateAccessToken(&user.User{
		ID:        parsedID,
		Email:     sessionUser.Email,
		FirstName: sessionUser.FirstName,
		LastName:  sessionUser.LastName,
	})
	if err != nil {
		return nil, err
	}

	return s.authRepo.UpdateAccessToken(accessToken, request.RefreshToken, parsedID)
}

func (s authService) Logout(sessionUser *context.SessionUser) error {
	parsedID, err := uuid.Parse(sessionUser.UserID)
	if err != nil {
		return errors.New("invalid user id format")
	}

	return s.authRepo.Logout(parsedID)
}

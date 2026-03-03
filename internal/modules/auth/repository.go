package auth

import (
	"errors"
	"sentinel-apm-api/internal/modules/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user *user.User) (*uuid.UUID, error)
	CheckUserEmail(email string) (*user.User, error)
	CreateRefreshToken(accessToken string, refreshToken string, userId uuid.UUID) (*RefreshTokenResponse, error)
	UpdateAccessToken(accessToken string, refreshToken string, userId uuid.UUID) (*RefreshTokenResponse, error)
	Logout(userId uuid.UUID) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) Register(user *user.User) (*uuid.UUID, error) {
	if err := a.db.Create(user).Error; err != nil {
		return nil, err
	}
	return &user.ID, nil
}

func (a *authRepository) CheckUserEmail(email string) (*user.User, error) {
	var user user.User
	if err := a.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *authRepository) CreateRefreshToken(accessToken string, refreshToken string, userId uuid.UUID) (*RefreshTokenResponse, error) {
	a.db.Create(&RefreshToken{
		ID:           uuid.New(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(time.Hour * 24 * 7),
		UserID:       userId,
	})

	return &RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    int(time.Now().Add(time.Hour * 24 * 7).Unix()),
	}, nil
}

func (a *authRepository) UpdateAccessToken(accessToken string, refreshToken string, userId uuid.UUID) (*RefreshTokenResponse, error) {

	var refreshTokenData RefreshToken
	if err := a.db.Where("access_token = ?", accessToken, "user_id = ?", userId).First(&refreshTokenData).Error; err != nil {
		return nil, err
	}

	if refreshTokenData.ExpiredAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	a.db.Model(&RefreshToken{}).Where("refresh_token = ?", refreshTokenData.RefreshToken, "user_id = ?", userId).Update("access_token", accessToken)

	return &RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    int(refreshTokenData.ExpiredAt.Unix()),
	}, nil
}

func (a *authRepository) Logout(userId uuid.UUID) error {
	if err := a.db.Where("user_id = ?", userId).Delete(&RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}

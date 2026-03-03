package auth

import (
	"sentinel-apm-api/internal/modules/user"
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID           string    `gorm:"type:varchar(255);primarykey;column:id"`
	UserID       uuid.UUID `gorm:"type:varchar(255);not null;column:user_id"`
	AccessToken  string    `gorm:"type:varchar(255);not null;column:access_token"`
	RefreshToken string    `gorm:"type:varchar(255);not null;column:refresh_token"`
	ExpiredAt    time.Time `gorm:"column:expired_at"`
	Device       string    `gorm:"type:varchar(255);not null;column:device"`
	IP           string    `gorm:"type:varchar(255);not null;column:ip"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`

	User user.User `gorm:"foreignKey:UserID;references:ID"`
}

type RefreshTokenRequest struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    int    `json:"expired_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int    `json:"expires_at"`
}

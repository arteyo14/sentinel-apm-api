package auth

import (
	"sentinel-apm-api/internal/modules/user"
	"time"
)

type RefreshToken struct {
	ID           string    `gorm:"type:varchar(255);primarykey;column:id"`
	UserID       string    `gorm:"type:varchar(255);not null;column:user_id"`
	AccessToken  string    `gorm:"type:varchar(255);not null;column:access_token"`
	RefreshToken string    `gorm:"type:varchar(255);not null;column:refresh_token"`
	ExpiresAt    time.Time `gorm:"column:expires_at"`
	Device       string    `gorm:"type:varchar(255);not null;column:device"`
	IP           string    `gorm:"type:varchar(255);not null;column:ip"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`

	User user.User `gorm:"foreignKey:UserID;references:ID"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int    `json:"expires_at"`
}

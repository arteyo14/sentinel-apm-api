package auth

import (
	"sentinel-apm-api/internal/modules/user"
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID           uuid.UUID `gorm:"type:varchar(255);primarykey;column:id"`
	UserID       uuid.UUID `gorm:"type:varchar(255);not null;column:user_id"`
	AccessToken  string    `gorm:"type:varchar(500);not null;column:access_token"`
	RefreshToken string    `gorm:"type:varchar(500);not null;column:refresh_token"`
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
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int    `json:"expires_at"`
}

type RegisterRequest struct {
	Email           string     `json:"email" binding:"required"`
	Password        string     `json:"password" binding:"required"`
	ConfirmPassword string     `json:"confirm_password" binding:"required"`
	FirstName       string     `json:"first_name" binding:"required"`
	LastName        string     `json:"last_name" binding:"required"`
	BirthDate       *time.Time `json:"birth_date" binding:"required"`
	Gender          string     `json:"gender" binding:"required"`
	Company         string     `json:"company" binding:"required"`
	RoleId          uuid.UUID  `json:"role_Id" binding:"required"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
}

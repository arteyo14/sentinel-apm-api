package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primarykey;column:id;default:uuid_generate_v4()"`
	FirstName string     `gorm:"type:varchar(255);not null;column:first_name"`
	LastName  string     `gorm:"type:varchar(255);not null;column:last_name"`
	Email     string     `gorm:"type:varchar(255);not null;unique;column:email"`
	Password  string     `gorm:"type:varchar(255);not null;column:password"`
	BirthDate *time.Time `gorm:"column:birth_date"`
	Gender    string     `gorm:"type:varchar(255);not null;column:gender"`
	Company   string     `gorm:"type:varchar(255);not null;column:company"`
	RoleID    uuid.UUID  `gorm:"type:uuid;not null;column:role_id"`
	Role      string     `gorm:"type:varchar(255);not null;column:role"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	IsEnabled bool       `gorm:"column:is_enabled"`
	IsActive  bool       `gorm:"column:is_active"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	ConfirmPassword string     `json:"confirm_password"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	BirthDate       *time.Time `json:"birth_date"`
	Gender          string     `json:"gender"`
	Company         string     `json:"company"`
	RoleId          uuid.UUID  `json:"role_Id"`
}

type UserResponse struct {
	ID        uuid.UUID  `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	BirthDate *time.Time `json:"birth_date"`
	Gender    string     `json:"gender"`
	Company   string     `json:"company"`
	RoleID    uuid.UUID  `json:"role_Id"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsEnabled bool       `json:"is_enabled"`
	IsActive  bool       `json:"is_active"`
}

package role

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID `gorm:"type:varchar(255);primarykey;column:id"`
	Name      string    `gorm:"type:varchar(255);not null;column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null;column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy uuid.UUID `gorm:"type:uuid;not null;column:updated_by"`
}

type RoleRequest struct {
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"created_by"`
}

type RoleResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`
}

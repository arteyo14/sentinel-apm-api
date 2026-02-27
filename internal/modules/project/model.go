package project

import (
	"sentinel-apm-api/internal/modules/user"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          string    `gorm:"type:varchar(255);primarykey;column:id"`
	Name        string    `gorm:"type:varchar(255);not null;column:name"`
	Description string    `gorm:"type:text;column:description"`
	StatusId    uuid.UUID `gorm:"type:uuid;not null;column:status_id"`
	Status      string    `gorm:"type:varchar(255);not null;column:status"`
	StatusColor string    `gorm:"type:varchar(255);not null;column:status_color"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null;column:created_by"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	UpdatedBy   uuid.UUID `gorm:"type:uuid;not null;column:updated_by"`

	ProjectMember []ProjectMember `gorm:"foreignKey:ProjectId;references:ID"`
}

type ProjectMember struct {
	ID         string    `gorm:"type:varchar(255);primarykey;column:id"`
	ProjectId  string    `gorm:"type:varchar(255);not null;column:project_id"`
	UserId     string    `gorm:"type:varchar(255);not null;column:user_id"`
	Role       string    `gorm:"type:varchar(255);not null;column:role"`
	MemberRole uuid.UUID `gorm:"type:uuid;not null;column:member_role"`
	Permission string    `gorm:"type:varchar(255);not null;column:permission"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	CreatedBy  uuid.UUID `gorm:"type:uuid;not null;column:created_by"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	UpdatedBy  uuid.UUID `gorm:"type:uuid;not null;column:updated_by"`

	Project Project   `gorm:"foreignKey:ProjectId;references:ID"`
	User    user.User `gorm:"foreignKey:UserId;references:ID"`
}

package memberrole

import (
	"sentinel-apm-api/internal/modules/project"
	"sentinel-apm-api/internal/modules/user"
	"time"

	"github.com/google/uuid"
)

type MemberRole struct {
	ID        string    `gorm:"type:varchar(255);primarykey;column:id"`
	ProjectId string    `gorm:"type:varchar(255);not null;column:project_id"`
	UserId    string    `gorm:"type:varchar(255);not null;column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null;column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy uuid.UUID `gorm:"type:uuid;not null;column:updated_by"`

	Project project.Project `gorm:"foreignKey:ProjectId;references:ID"`
	User    user.User       `gorm:"foreignKey:UserId;references:ID"`
}

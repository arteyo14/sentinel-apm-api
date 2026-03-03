package role

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *Role) (*uuid.UUID, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) CreateRole(role *Role) (*uuid.UUID, error) {
	if err := r.db.Create(role).Error; err != nil {
		return nil, err
	}
	return &role.ID, nil
}

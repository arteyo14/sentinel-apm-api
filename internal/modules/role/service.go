package role

import "github.com/google/uuid"

type RoleService interface {
	CreateRole(role RoleRequest) (*uuid.UUID, error)
}

type roleService struct {
	roleRepo RoleRepository
}

func NewRoleService(roleRepo RoleRepository) RoleService {
	return &roleService{roleRepo: roleRepo}
}

func (r *roleService) CreateRole(role RoleRequest) (*uuid.UUID, error) {
	var roleData = Role{
		Name:      role.Name,
		CreatedBy: role.CreatedBy,
	}

	return r.roleRepo.CreateRole(&roleData)
}

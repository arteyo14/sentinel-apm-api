package migrate

import (
	"fmt"
	"sentinel-apm-api/internal/modules/auth"
	memberrole "sentinel-apm-api/internal/modules/member-role"
	"sentinel-apm-api/internal/modules/project"
	"sentinel-apm-api/internal/modules/role"
	"sentinel-apm-api/internal/modules/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&user.User{},
		&project.Project{},
		&project.ProjectMember{},
		&memberrole.MemberRole{},
		&auth.RefreshToken{},
		&role.Role{},
	); err != nil {
		panic("failed to migrate database")
	}

	fmt.Println("✅ Database migrated successfully!")
}

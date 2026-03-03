package seed

import (
	"log"
	"sentinel-apm-api/internal/modules/role"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	var role role.Role
	role.ID = uuid.New()
	role.Name = "Super Admin"
	role.IsActive = true
	role.IsEnabled = true
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()

	if err := db.Create(&role).Error; err != nil {
		panic(err)
	}

	log.Println("Role seeded successfully")
}

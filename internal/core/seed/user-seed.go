package seed

import (
	"log"
	"sentinel-apm-api/internal/modules/role"
	"sentinel-apm-api/internal/modules/user"
	"sentinel-apm-api/utils/hash"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	var role role.Role

	db.Where("name = ?", "Super Admin").First(&role)

	var user user.User
	user.ID = uuid.New()
	user.Email = "admin@admin.com"
	user.Password = hash.HashPassword("P@ssw0rd!")
	user.FirstName = "Super"
	user.LastName = "Admin"
	user.IsActive = true
	user.Gender = "Male"
	user.Company = "Sentinel APM"
	user.BirthDate = time.Date(1998, 9, 25, 0, 0, 0, 0, time.UTC)
	user.RoleID = uuid.New()
	user.Role = "Super Admin"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := db.Create(&user).Error; err != nil {
		panic(err)
	}

	log.Println("User seeded successfully")
}

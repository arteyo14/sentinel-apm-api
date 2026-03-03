package seed

import (
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	SeedRole(db)
	SeedUser(db)
}

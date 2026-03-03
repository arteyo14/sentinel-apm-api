package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) (*uuid.UUID, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Create(user *User) (*uuid.UUID, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}
	return &user.ID, nil
}

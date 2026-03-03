package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) (*uuid.UUID, error)
	GetUsers() ([]UserResponse, error)
	UpdateUser(userId uuid.UUID, user *User) error
	DeleteUser(userId uuid.UUID) error
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

func (repo *userRepository) GetUsers() ([]UserResponse, error) {
	var users []User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			BirthDate: user.BirthDate,
			Gender:    user.Gender,
			Company:   user.Company,
			RoleID:    user.RoleID,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			IsEnabled: user.IsEnabled,
			IsActive:  user.IsActive,
		})
	}

	return userResponses, nil
}

func (repo *userRepository) UpdateUser(userId uuid.UUID, user *User) error {
	if err := repo.db.Where("id = ?", userId).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) DeleteUser(userId uuid.UUID) error {
	if err := repo.db.Where("id = ?", userId).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

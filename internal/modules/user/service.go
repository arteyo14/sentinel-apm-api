package user

import (
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req *UserRequest) (*uuid.UUID, error)
	GetUsers() ([]UserResponse, error)
	UpdateUser(userId uuid.UUID, req *UpdateUserRequest) error
	DeleteUser(userId uuid.UUID) error
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(req *UserRequest) (userId *uuid.UUID, err error) {
	var birthdate *time.Time
	if req.BirthDate != nil {
		birthdate = req.BirthDate
	}

	user := &User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		BirthDate: birthdate,
		Gender:    req.Gender,
		Company:   req.Company,
		RoleID:    req.RoleId,
	}

	userId, err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return userId, nil
}

func (s *userService) GetUsers() ([]UserResponse, error) {
	users, err := s.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) UpdateUser(userId uuid.UUID, req *UpdateUserRequest) error {
	user := &User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		BirthDate: req.BirthDate,
		Gender:    req.Gender,
		Company:   req.Company,
		RoleID:    req.RoleID,
		IsActive:  req.IsActive,
	}

	if err := s.userRepo.UpdateUser(userId, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(userId uuid.UUID) error {
	if err := s.userRepo.DeleteUser(userId); err != nil {
		return err
	}
	return nil
}

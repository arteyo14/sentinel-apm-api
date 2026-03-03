package user

import (
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req *UserRequest) (*uuid.UUID, error)
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
		BirthDate: *birthdate,
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

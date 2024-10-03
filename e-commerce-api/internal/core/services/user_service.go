package services

import (
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
)

type userService struct {
	userRepo ports.UserRepository
}

// CreateUser implements ports.UserService.
func (u *userService) CreateUser(createUserReq dtos.UserSignUpRequest) (domain.User, error) {
	panic("unimplemented")
}

// GetUser implements ports.UserService.
func (u *userService) GetUser(userID string) (domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements ports.UserService.
func (u *userService) UpdateUser(userID string, user domain.User) (domain.User, error) {
	panic("unimplemented")
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

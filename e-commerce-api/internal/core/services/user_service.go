package services

import (
	"errors"
	"log"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/common/utils"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
)

type userService struct {
	userRepo ports.UserRepository
}

// CreateUser implements ports.UserService.
func (u *userService) CreateUser(createUserReq dtos.UserSignUpRequest) (domain.User, error) {
	// check if user is already present or not by checking email
	user, err := u.userRepo.GetUserByEmail(createUserReq.Email)
	if err != nil {
		log.Println("UserService::CreateUser: creating new user")

		hashedPassword, err := utils.HashPassword(createUserReq.Password)
		if err != nil {
			// return error while hashing password
			return domain.User{}, err
		}

		newUser := domain.User{
			FirstName: createUserReq.FirstName,
			LastName: createUserReq.LastName,
			UserEmail: createUserReq.Email,
			UserPassword: hashedPassword,
		}

		// call repo to create new user
		user, err = u.userRepo.CreateUser(newUser)
		if err != nil {
			return domain.User{}, err
		}

		// return the new user created
		return user, nil
	}

	log.Println("UserService::CreateUser: email already taken: ", user.UserEmail)

	return domain.User{}, errors.New("email already taken");
}

// GetUser implements ports.UserService.
func (u *userService) GetUser(loginReq dtos.UserLoginRequest) (domain.User, error) {
	user, err := u.userRepo.GetUserByEmail(loginReq.Email)
	if err != nil {
		// user not found return error
		return domain.User{}, errors.New("user not registered with the following email")
	}

	// check password
	validPassword := utils.CheckPasswordHash(loginReq.Password, user.UserPassword)
	if !validPassword {
		return domain.User{}, errors.New("invalid password")
	}

	return user, nil
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

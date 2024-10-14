package services

import (
	"errors"
	"log"

	"github.com/adityasunny1189/roadmap-sh/auth/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/auth/internal/core/ports"
)

type authService struct {
	authRepo ports.AuthRepository
}

func NewAuthService(authRepo ports.AuthRepository) ports.AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (a *authService) SignIn(email, password string) (string, error) {
	user, err := a.authRepo.GetUserByEmail(email)
	if err != nil {
		log.Println("error while getting user: ", err)
		return "", err
	}

	if user.CheckPasswordHash(user.Password, password) {
		// TODO: Create a new bearer token and send it
		return "bearerToken", nil
	} else {
		log.Println("Invalid Password")
		return "", errors.New("incorrect password")
	}
}

func (a *authService) SignUp(user domain.User) (string, error) {
	// check for duplicate details
	_, err := a.authRepo.GetUserByEmail(user.Email)
	if err != nil {
		u, err := a.authRepo.CreateUser(user)
		if err != nil {
			log.Println("error while creating user: ", err)
			return "", err
		}

		// TODO: generate new bearer token
		return u.Email, nil
	}
	return "", errors.New("user already present with the given details")
}

func (a *authService) Verify(bearerToken string) (bool, domain.User) {
	return true, domain.User{}
}

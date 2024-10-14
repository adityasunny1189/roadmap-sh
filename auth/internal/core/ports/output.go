package ports

import "github.com/adityasunny1189/roadmap-sh/auth/internal/core/domain"

type AuthService interface {
	SignIn(email, password string) (string, error)
	SignUp(user domain.User) (string, error)
	Verify(bearerToken string) (bool, domain.User)
}

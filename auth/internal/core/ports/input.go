package ports

import "github.com/adityasunny1189/roadmap-sh/auth/internal/core/domain"

type AuthRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	GetUserByUsername(username string) (domain.User, error)
}

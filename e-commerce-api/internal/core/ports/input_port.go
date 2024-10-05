package ports

import (
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUser(userId string) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	UpdateUser(user domain.User) (domain.User, error)
}

type ProductRepository interface {
}

type CartRepository interface {
}

type OrderRepository interface {
}

type PaymentRepository interface {
}

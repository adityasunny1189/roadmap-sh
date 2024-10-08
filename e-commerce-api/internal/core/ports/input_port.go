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
	CreateProduct(product domain.Product) (domain.Product, error)
	GetProduct(productId string) (domain.Product, error)
	GetAllProductByCategory(categoryName string) ([]domain.Product, error)
	SearchProduct(keyword string) ([]domain.Product, error)
	GetAllProducts() ([]domain.Product, error)	
}

type CartRepository interface {
}

type CheckoutRepository interface {
}

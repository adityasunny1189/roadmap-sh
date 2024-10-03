package ports

import (
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
)

type UserService interface {
	CreateUser(createUserReq dtos.UserSignUpRequest) (domain.User, error)
	GetUser(userID string) (domain.User, error)
	UpdateUser(userID string, user domain.User) (domain.User, error)
}

type ProductService interface {
	AddProduct(product domain.Product) (domain.Product, error)
	GetProduct(productID string) (domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	UpdateProductStock(productID string, quantity int) (domain.Product, error)
}

type CartService interface {
	CreateCart(cart domain.Cart) (domain.Cart, error)
	GetCart(cartID string) (domain.Cart, error)
	AddItemToCart(cartItem domain.CartItem) (domain.CartItem, error)
	RemoveItemFromCart(cartItemID string) error
	UpdateCart(cart domain.Cart) (domain.Cart, error)
	CheckoutCart(cartID string) error
}

type OrderService interface {
	CreateOrder(order domain.Order) (domain.Order, error)
	GetOrder(orderID string) (domain.Order, error)
	GetOrdersByUserID(userID string) ([]domain.Order, error)
}

type PaymentService interface {
}

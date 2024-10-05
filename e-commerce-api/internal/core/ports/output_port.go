package ports

import (
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
)

type UserService interface {
	CreateUser(createUserReq dtos.UserSignUpRequest) (domain.User, error)
	GetUser(loginReq dtos.UserLoginRequest) (domain.User, error)
	UpdateUser(userID string, user domain.User) (domain.User, error)
}

type ProductService interface {
	AddProduct(addProductReq dtos.AddNewProductRequest) (domain.Product, error)
	GetProductById(productID string) (domain.Product, error)
	GetProductsByCategory(categoryName string) ([]domain.Product, error)
	GetProductsByKeyword(keyword string) ([]domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	UpdateProductStock(updateProductInventory dtos.UpdateInventoryRequest) (int64, error)
	SortAndFilterProduct(sortAndFilterReq dtos.SortAndFilterProductRequest) ([]domain.Product, error)
}

type CartService interface {
	GetAllCarts() ([]domain.Cart, error)
	GetCartDetails(cartID string) (domain.Cart, []domain.CartItem, error)
	CreateCart(createCartReq dtos.CreateCartRequest) (domain.Cart, error)
	UpdateCart(updateCartReq dtos.UpdateCartRequest) (domain.Cart, []domain.CartItem, error)
	DeleteCart(cartID string) (domain.Cart, error)
}

type CheckoutService interface {
	CreateOrder(createOrderReq dtos.CreateOrderRequest) (domain.Order, error)
	GetAllOrders() ([]domain.Order, error)
	GetOrderDetails(orderID string) (domain.Order, []domain.CartItem, error)
	GetOrderStatus(orderId string) (domain.Order, error)
	InititatePayment(paymentReq dtos.PaymentRequest) (domain.Payment, error)
}

package services

import (
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
)

type checkoutService struct {
	checkoutRepo ports.CheckoutRepository
}

// CreateOrder implements ports.CheckoutService.
func (*checkoutService) CreateOrder(createOrderReq dtos.CreateOrderRequest) (domain.Order, error) {
	panic("unimplemented")
}

// GetAllOrders implements ports.CheckoutService.
func (*checkoutService) GetAllOrders() ([]domain.Order, error) {
	panic("unimplemented")
}

// GetOrderDetails implements ports.CheckoutService.
func (*checkoutService) GetOrderDetails(orderID string) (domain.Order, []domain.CartItem, error) {
	panic("unimplemented")
}

// GetOrderStatus implements ports.CheckoutService.
func (*checkoutService) GetOrderStatus(orderId string) (domain.Order, error) {
	panic("unimplemented")
}

// InititatePayment implements ports.CheckoutService.
func (*checkoutService) InititatePayment(paymentReq dtos.PaymentRequest) (domain.Payment, error) {
	panic("unimplemented")
}

func NewCheckoutService(checkoutRepo ports.CheckoutRepository) ports.CheckoutService {
	return &checkoutService{
		checkoutRepo: checkoutRepo,
	}
}

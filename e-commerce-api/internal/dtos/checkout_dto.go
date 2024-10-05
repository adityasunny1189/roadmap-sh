package dtos

import "github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"

type GetAllOrdersResponse struct {
	Orders []domain.Order
}

type CreateOrderRequest struct {
	CartId string
}

type CreateOrderResponse struct {
	Order domain.Order
}

type GetOrderDetailsResponse struct {
	OrderDetails domain.Order
	Items        []domain.CartItem
}

type OrderStatusPollingResponse struct {
	Status string
}

type PaymentRequest struct {
	OrderId       string
	PaymentMethod string
}

type PaymentResponse struct {
	Status string
}

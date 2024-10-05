package dtos

import "github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"

type CartDto struct {
	CartMetadata domain.Cart
	CartItems    []domain.CartItem
}

type GetAllCartsResponse struct {
	Cart []domain.Cart
}

type GetCartDetailsResponse struct {
	CartDetails CartDto
}

type CreateCartRequest struct {
	ProductId       string
	ProductPrice    float64
	ProductQuantity int64
}

type CreateCartResponse struct {
	Cart domain.Cart
}

type UpdateCartRequest struct {
	ProductId       string
	ProductPrice    float64
	ProductQuantity int64
}

type UpdateCartResponse struct {
	CartDetails CartDto
}

type DeleteCartResponse struct {
	Cart domain.Cart
}

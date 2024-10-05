package services

import (
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
)

type cartService struct {
	cartRepo ports.CartRepository
}

// CreateCart implements ports.CartService.
func (*cartService) CreateCart(createCartReq dtos.CreateCartRequest) (domain.Cart, error) {
	panic("unimplemented")
}

// DeleteCart implements ports.CartService.
func (*cartService) DeleteCart(cartID string) (domain.Cart, error) {
	panic("unimplemented")
}

// GetAllCarts implements ports.CartService.
func (*cartService) GetAllCarts() ([]domain.Cart, error) {
	panic("unimplemented")
}

// GetCartDetails implements ports.CartService.
func (*cartService) GetCartDetails(cartID string) (domain.Cart, []domain.CartItem, error) {
	panic("unimplemented")
}

// UpdateCart implements ports.CartService.
func (*cartService) UpdateCart(updateCartReq dtos.UpdateCartRequest) (domain.Cart, []domain.CartItem, error) {
	panic("unimplemented")
}

func NewCartService(cartRepo ports.CartRepository) ports.CartService {
	return &cartService{
		cartRepo: cartRepo,
	}
}

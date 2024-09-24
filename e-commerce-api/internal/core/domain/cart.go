package domain

import (
	"github.com/google/uuid"
)

type CartState string

const (
	CREATED   CartState = "CART_CREATED"
	COMPLETED CartState = "CART_COMPLETED"
	CANCELED  CartState = "CART_CANCELED"
)

type Cart struct {
	Id         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"user_id"`
	CartState  CartState `json:"cart_state"`
	CartAmount float64   `json:"cart_amount"`
}

type CartItem struct {
	Id        uuid.UUID `json:"id"`
	CartId    uuid.UUID `json:"cart_id"`
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

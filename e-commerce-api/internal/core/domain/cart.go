package domain

import (
	"github.com/google/uuid"
)

type CartState string

const (
	CART_CREATED   CartState = "CART_CREATED"
	CART_COMPLETED CartState = "CART_COMPLETED"
	CART_CANCELED  CartState = "CART_CANCELED"
)

type Cart struct {
	Id         uuid.UUID `json:"id"`
	UserId     int64     `json:"user_id"`
	CartState  CartState `json:"cart_state"`
	CartAmount float64   `json:"cart_amount"`
}

type CartItem struct {
	Id        uuid.UUID `json:"id"`
	CartId    uuid.UUID `json:"cart_id"`
	ProductId int64     `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

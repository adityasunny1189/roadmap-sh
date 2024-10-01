package domain

import (
	"github.com/google/uuid"
)

type OrderState string

const (
	ORDER_CREATED   OrderState = "ORDER_CREATED"
	ORDER_COMPLETED OrderState = "ORDER_COMPLETED"
	ORDER_CANCELLED OrderState = "ORDER_CANCELLED"
)

type Order struct {
	Id         int64      `json:"id"`
	UserID     int64      `json:"user_id"`
	CartId     uuid.UUID  `json:"cart_id"`
	TotalPrice float64    `json:"total_price"`
	OrderState OrderState `json:"order_state"`
}

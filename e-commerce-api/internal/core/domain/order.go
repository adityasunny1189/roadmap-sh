package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderState string

const (
	CREATED   OrderState = "ORDER_CREATED"
	COMPLETED OrderState = "ORDER_COMPLETED"
	CANCELLED OrderState = "ORDER_CANCELLED"
)

type Order struct {
	Id         uuid.UUID  `json:"id"`
	UserID     uuid.UUID  `json:"user_id"`
	CartId     uuid.UUID  `json:"cart_id"`
	TotalPrice float64    `json:"total_price"`
	OrderState OrderState `json:"order_state"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

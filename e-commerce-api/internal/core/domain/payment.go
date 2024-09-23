package domain

import (
	"time"

	"github.com/google/uuid"
)

type PaymentState string

const (
	PENDING  PaymentState = "PAYMENT_PENDING"
	SUCCESS  PaymentState = "PAYMENT_SUCCESS"
	FAILED   PaymentState = "PAYMENT_FAILED"
	CANCELED PaymentState = "PAYMENT_CANCELED"
)

type PaymentMethod string

const (
	RAZORPAY PaymentMethod = "RAZORPAY"
	STRIPE   PaymentMethod = "STRIPE"
	PAYPAL   PaymentMethod = "PAYPAL"
)

type Payment struct {
	Id            uuid.UUID     `json:"id"`
	OrderId       uuid.UUID     `json:"order_id"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	PaymentState  PaymentState  `json:"payment_state"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

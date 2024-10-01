package domain

type PaymentState string

const (
	PAYMENT_PENDING  PaymentState = "PAYMENT_PENDING"
	PAYMENT_SUCCESS  PaymentState = "PAYMENT_SUCCESS"
	PAYMENT_FAILED   PaymentState = "PAYMENT_FAILED"
	PAYMENT_CANCELED PaymentState = "PAYMENT_CANCELED"
)

type PaymentMethod string

const (
	RAZORPAY PaymentMethod = "RAZORPAY"
	STRIPE   PaymentMethod = "STRIPE"
	PAYPAL   PaymentMethod = "PAYPAL"
)

type Payment struct {
	Id            int64         `json:"id"`
	OrderId       int64         `json:"order_id"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	PaymentState  PaymentState  `json:"payment_state"`
}

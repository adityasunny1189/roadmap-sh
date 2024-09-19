package domain

import "github.com/google/uuid"

type PaymentState string

// Todo: Create all possible payment state
const (
	SUCCESS PaymentState = "Success"
	FAILED  PaymentState = "Failed"
)

// Todo: Create state machine to handle payment state change

type Payment struct {
	Id      uuid.UUID
	UserId  uuid.UUID
	OrderId uuid.UUID
	State   PaymentState
	Amount  float64
}

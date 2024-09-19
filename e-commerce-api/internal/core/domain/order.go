package domain

import "github.com/google/uuid"

type OrderState string

const (
	CREATED OrderState = "Created"
	COMPLETED OrderState = ""
	PAYMENT_FAILURE OrderState = ""
	CANCELLED OrderState = ""
	// Todo: To add more state 
)

// Todo: Create a state machine to handle state change

type Order struct {
	Id     uuid.UUID
	CartId uuid.UUID
	State  OrderState
}

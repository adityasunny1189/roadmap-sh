package domain

import "github.com/google/uuid"

type Inventory struct {
	ProductQuantity map[uuid.UUID]int
}

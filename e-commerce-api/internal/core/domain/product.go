package domain

import "github.com/google/uuid"

type Product struct {
	ID                 uuid.UUID `json:"id"`
	ProductName        string    `json:"product_name"`
	ProductDescription string    `json:"product_description"`
	Category           string    `json:"category"`
	ImageURL           string    `json:"image_url"`
	Price              float64   `json:"price"`
}

type Inventory map[uuid.UUID]int

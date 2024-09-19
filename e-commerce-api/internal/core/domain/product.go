package domain

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	ImageURL    string    `json:"image_url"`
	Price       float64   `json:"price"`
}

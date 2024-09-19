package domain

import "github.com/google/uuid"

type Cart struct {
	Id          uuid.UUID         `json:"id"`
	UserId      uuid.UUID         `json:"user_id"`
	ProductList map[uuid.UUID]int `json:"product_list"`
	TotalAmount float64
}

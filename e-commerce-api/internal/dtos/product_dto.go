package dtos

import "github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"

type AddNewProductRequest struct {
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	CategoryId         int64   `json:"category_id"`
	ImageURL           string  `json:"image_url"`
	Price              float64 `json:"price"`
}

type AddNewProductResponse struct {
	Product domain.Product `json:"product"`
}

type UpdateInventoryRequest struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type UpdateInventoryResponse struct {
	CurrentQuantity int64 `json:"current_quantity"`
}

type SortAndFilterOption string

const (
	SORT_BY_PRICE    SortAndFilterOption = "SORT_BY_PRICE"
	SORT_BY_CATEGORY SortAndFilterOption = "SORT_BY_CATEGORY"
)

type SortAndFilterProductRequest struct {
	SortByOption SortAndFilterOption `json:"sort_by_option"`
}

type ProductsResponse struct {
	Products []domain.Product `json:"products"`
}

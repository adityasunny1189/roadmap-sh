package domain

type Category struct {
	Id                  int64  `json:"id"`
	CategoryName        string `json:"category_name"`
	CategoryDescription string `json:"category_description"`
}

type Product struct {
	ID                 int64   `json:"id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	CategoryId         int64   `json:"category_id"`
	ImageURL           string  `json:"image_url"`
	Price              float64 `json:"price"`
}

type Inventory struct {
	Id        int64 `json:"id"`
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

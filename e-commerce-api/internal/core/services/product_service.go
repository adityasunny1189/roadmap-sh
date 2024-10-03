package services

import (
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
)

type productService struct {
	productRepo ports.ProductRepository
}

// SortAndFilterProduct implements ports.ProductService.
func (*productService) SortAndFilterProduct(sortAndFilterReq dtos.SortAndFilterProductRequest) ([]domain.Product, error) {
	panic("unimplemented")
}

// UpdateProductStock implements ports.ProductService.
func (*productService) UpdateProductStock(updateProductInventory dtos.UpdateInventoryRequest) (int64, error) {
	panic("unimplemented")
}

// GetProductsByKeyword implements ports.ProductService.
func (*productService) GetProductsByKeyword(keyword string) ([]domain.Product, error) {
	panic("unimplemented")
}

// GetProductByCategory implements ports.ProductService.
func (*productService) GetProductsByCategory(categoryName string) ([]domain.Product, error) {
	panic("unimplemented")
}

// AddProduct implements ports.ProductService.
func (p *productService) AddProduct(addProductReq dtos.AddNewProductRequest) (domain.Product, error) {
	panic("unimplemented")
}

// GetAllProducts implements ports.ProductService.
func (p *productService) GetAllProducts() ([]domain.Product, error) {
	panic("unimplemented")
}

// GetProduct implements ports.ProductService.
func (p *productService) GetProductById(productID string) (domain.Product, error) {
	panic("unimplemented")
}

func NewProductService(productRepo ports.ProductRepository) ports.ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

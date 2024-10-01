package services

import (
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
)

type productService struct {
	productRepo ports.ProductRepository
}

// AddProduct implements ports.ProductService.
func (p *productService) AddProduct(product domain.Product) (domain.Product, error) {
	panic("unimplemented")
}

// GetAllProducts implements ports.ProductService.
func (p *productService) GetAllProducts() ([]domain.Product, error) {
	panic("unimplemented")
}

// GetProduct implements ports.ProductService.
func (p *productService) GetProduct(productID string) (domain.Product, error) {
	panic("unimplemented")
}

// UpdateProductStock implements ports.ProductService.
func (p *productService) UpdateProductStock(productID string, quantity int) (domain.Product, error) {
	panic("unimplemented")
}

func NewProductService(productRepo ports.ProductRepository) ports.ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

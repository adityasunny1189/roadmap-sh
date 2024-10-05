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
func (p *productService) SortAndFilterProduct(sortAndFilterReq dtos.SortAndFilterProductRequest) ([]domain.Product, error) {
	panic("unimplemented")
}

// UpdateProductStock implements ports.ProductService.
func (p *productService) UpdateProductStock(updateProductInventory dtos.UpdateInventoryRequest) (int64, error) {
	panic("unimplemented")
}

// GetProductsByKeyword implements ports.ProductService.
func (p *productService) GetProductsByKeyword(keyword string) ([]domain.Product, error) {
	products, err := p.productRepo.SearchProduct(keyword)
	if err != nil {
		return []domain.Product{}, err
	}

	return products, nil
}

// GetProductByCategory implements ports.ProductService.
func (p *productService) GetProductsByCategory(categoryName string) ([]domain.Product, error) {
	products, err := p.productRepo.GetAllProductByCategory(categoryName)
	if err != nil {
		return []domain.Product{}, err
	}

	return products, nil
}

// AddProduct implements ports.ProductService.
func (p *productService) AddProduct(addProductReq dtos.AddNewProductRequest) (domain.Product, error) {
	newProduct := domain.Product{
		ProductName:        addProductReq.ProductName,
		ProductDescription: addProductReq.ProductDescription,
		CategoryId:         addProductReq.CategoryId,
		ImageURL:           addProductReq.ImageURL,
		Price:              addProductReq.Price,
	}

	product, err := p.productRepo.CreateProduct(newProduct)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

// GetAllProducts implements ports.ProductService.
func (p *productService) GetAllProducts() ([]domain.Product, error) {
	products, err := p.productRepo.GetAllProducts()
	if err != nil {
		return []domain.Product{}, err
	}

	return products, nil
}

// GetProduct implements ports.ProductService.
func (p *productService) GetProductById(productID string) (domain.Product, error) {
	product, err := p.productRepo.GetProduct(productID)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func NewProductService(productRepo ports.ProductRepository) ports.ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

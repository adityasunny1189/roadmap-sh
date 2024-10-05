package repository

import (
	"database/sql"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
)

type productRepositoryImpl struct {
	db *sql.DB
}

// CreateProduct implements ports.ProductRepository.
func (*productRepositoryImpl) CreateProduct(product domain.Product) (domain.Product, error) {
	panic("unimplemented")
}

// GetAllProductByCategory implements ports.ProductRepository.
func (*productRepositoryImpl) GetAllProductByCategory(categoryName string) ([]domain.Product, error) {
	panic("unimplemented")
}

// GetAllProducts implements ports.ProductRepository.
func (*productRepositoryImpl) GetAllProducts() ([]domain.Product, error) {
	panic("unimplemented")
}

// GetProduct implements ports.ProductRepository.
func (*productRepositoryImpl) GetProduct(productId string) (domain.Product, error) {
	panic("unimplemented")
}

// SearchProduct implements ports.ProductRepository.
func (*productRepositoryImpl) SearchProduct(keyword string) ([]domain.Product, error) {
	panic("unimplemented")
}

func NewProductRepository(db *sql.DB) ports.ProductRepository {
	return &productRepositoryImpl{
		db: db,
	}
}

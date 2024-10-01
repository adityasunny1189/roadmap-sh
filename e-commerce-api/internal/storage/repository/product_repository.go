package repository

import (
	"database/sql"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
)

type productRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ports.ProductRepository {
	return &productRepositoryImpl{
		db: db,
	}
}

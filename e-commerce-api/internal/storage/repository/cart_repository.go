package repository

import (
	"database/sql"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
)

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) ports.CartRepository {
	return &cartRepository{
		db: db,
	}
}

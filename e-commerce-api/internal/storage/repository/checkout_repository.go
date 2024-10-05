package repository

import (
	"database/sql"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
)

type checkoutRepository struct {
	db *sql.DB
}

func NewCheckoutRepository(db *sql.DB) ports.CartRepository {
	return &checkoutRepository{
		db: db,
	}
}

package repository

import (
	"database/sql"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
)

type userRepositoryImpl struct {
	db *sql.DB
}

// CreateUser implements ports.UserRepository.
func (u *userRepositoryImpl) CreateUser(user domain.User) (domain.User, error) {
	panic("unimplemented")
}

// GetUser implements ports.UserRepository.
func (u *userRepositoryImpl) GetUser(userId string) (domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements ports.UserRepository.
func (u *userRepositoryImpl) UpdateUser(user domain.User) (domain.User, error) {
	panic("unimplemented")
}

func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

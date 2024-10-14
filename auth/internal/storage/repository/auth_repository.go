package repository

import (
	"database/sql"

	"github.com/adityasunny1189/roadmap-sh/auth/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/auth/internal/core/ports"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

// CreateUser implements ports.AuthRepository.
func (a *authRepository) CreateUser(user domain.User) (domain.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements ports.AuthRepository.
func (a *authRepository) GetUserByEmail(email string) (domain.User, error) {
	panic("unimplemented")
}

// GetUserByUsername implements ports.AuthRepository.
func (a *authRepository) GetUserByUsername(username string) (domain.User, error) {
	panic("unimplemented")
}

func NewAuthRepository(conn *sql.DB) ports.AuthRepository {
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &authRepository{
		db: db,
	}
}

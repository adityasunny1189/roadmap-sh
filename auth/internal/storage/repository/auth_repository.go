package repository

import (
	"database/sql"
	"log"

	"github.com/adityasunny1189/roadmap-sh/auth/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/auth/internal/core/ports"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func (a *authRepository) CreateUser(user domain.User) (domain.User, error) {
	if err := a.db.Create(user).Error; err != nil {
		log.Println("Can't create user: ", err)
		return domain.User{}, err
	}

	return user, nil
}

func (a *authRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := a.db.First(&user, "email = ?", email).Error; err != nil {
		log.Println("Can't find data: ", err)
		return user, err
	}
	return user, nil
}

func (a *authRepository) GetUserByUsername(username string) (domain.User, error) {
	var user domain.User
	if err := a.db.First(&user, "username = ?", username).Error; err != nil {
		log.Println("Can't find data: ", err)
		return user, err
	}
	return user, nil
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

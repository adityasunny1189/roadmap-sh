package utils

import "golang.org/x/crypto/bcrypt"

const (
	SECRET_KEY string = "secret-key"
	ADMIN_EMAIL string = ""
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

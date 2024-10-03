package auth

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/common/utils"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateNewToken(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userEmail": user.UserEmail,
			"exp":       time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(utils.SECRET_KEY))
	if err != nil {
		log.Println("Error: ", err)
		return "", err
	}
	log.Println("Token: ", tokenString)
	return "Bearer " + tokenString, nil
}

func VerifyToken(tokenString string) (interface{}, error) {
	actualToken := strings.Split(tokenString, " ")
	if len(actualToken) != 2 {
		return nil, errors.New("empty bearer token")
	}

	token, err := jwt.Parse(actualToken[1], func(t *jwt.Token) (interface{}, error) {
		return []byte(utils.SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("INVALID TOKEN")
	}

	userEmail := token.Claims.(jwt.MapClaims)["userEmail"]

	return userEmail, nil
}

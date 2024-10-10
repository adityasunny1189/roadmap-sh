package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/common/utils"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
	"github.com/golang-jwt/jwt/v5"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(userRepo ports.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if authHeader == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			bearerToken := strings.Split(authHeader, " ")
			token, err := jwt.Parse(bearerToken[1], func(t *jwt.Token) (interface{}, error) {
				return []byte(utils.SECRET_KEY), nil
			})

			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			if !token.Valid {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			useremail := token.Claims.(jwt.MapClaims)["userEmail"].(string)

			// Find user from user repo
			user, err := userRepo.GetUserByEmail(useremail)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func GetUserFromContext(ctx context.Context) *domain.User {
	user, _ := ctx.Value(userCtxKey).(*domain.User)
	return user
}

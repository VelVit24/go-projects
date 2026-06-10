package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/VelVit24/todo-api/service"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		if len(tokenString) == 0 {
			http.Error(w, "not logined", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&service.Claims{},
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(os.Getenv("KEY_JWT")), nil
			},
		)
		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		id := token.Claims.(*service.Claims).UserID
		ctx := context.WithValue(r.Context(), service.UserIDKey, id)
		next(w, r.WithContext(ctx))
	}
}

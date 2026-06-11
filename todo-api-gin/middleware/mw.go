package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/VelVit24/todo-api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		if len(tokenString) == 0 {
			c.JSON(http.StatusUnauthorized, "not logined")
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
			c.JSON(http.StatusUnauthorized, "invalid token")
			return
		}
		id := token.Claims.(*service.Claims).UserID
		c.Set(service.UserIDKey, id)
		c.Next()
	}
}

var apiLimiter = rate.NewLimiter(rate.Every(5*time.Second), 1)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !apiLimiter.Allow() {
			c.JSON(http.StatusTooManyRequests, "Too Many Requests")
			return
		}
		c.Next()
	}
}

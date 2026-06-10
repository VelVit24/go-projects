package service

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"os"
	"time"

	"github.com/VelVit24/todo-api/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type contextKey string

const UserIDKey contextKey = "userID"

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func WriteJson(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}

func GenToken(id int) (string, error) {
	key := []byte(os.Getenv("KEY_JWT"))
	claims := Claims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func ValidateLogin(user models.User) bool {
	_, err := mail.ParseAddress(user.Email)
	if len(user.Password) < 8 || err != nil {
		return false
	}
	return true
}

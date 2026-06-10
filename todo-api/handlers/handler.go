package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/VelVit24/todo-api/database"
	"github.com/VelVit24/todo-api/models"
	"github.com/VelVit24/todo-api/service"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	DB *sql.DB
}

var key string = "key"

func (h *Handler) HandleReg(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	res, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(res, &user)
	if err != nil {
		log.Println(err)
		return
	}
	// валидация
	err = database.InsertUser(h.DB, &user)
	if err != nil {
		log.Println(err)
		return
	}
	tokenString, err := service.GenToken(user.Id)
	if err != nil {
		log.Println(err)
	}
	token := models.Token{Token: tokenString}
	service.WriteJson(w, 200, token)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	res, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(res, &user)
	if err != nil {
		log.Println(err)
		return
	}
	// валидация
	ok, err := database.CheckUser(h.DB, &user)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "incorect email or password", http.StatusUnauthorized)
			return
		}
		log.Println(err)
		return
	}
	if !ok {
		http.Error(w, "incorect email or password", http.StatusUnauthorized)
		return
	}
	tokenString, err := service.GenToken(user.Id)
	if err != nil {
		log.Println(err)
	}
	token := models.Token{Token: tokenString}
	service.WriteJson(w, 200, token)
}

func (h *Handler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	log.Println(tokenString)
	if len(tokenString) == 0 {
		http.Error(w, "not logined", http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		http.Error(w, "not logined", http.StatusUnauthorized)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id := int(claims["user_id"].(float64))

	log.Println(id)

	todo := models.Todo{}
	res, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(res, &todo)
	if err != nil {
		log.Println(err)
		return
	}
	// валидация

	err = database.CreateTodo(h.DB, id, &todo)
	if err != nil {
		log.Println(err)
		return
	}
	service.WriteJson(w, 200, todo)
}

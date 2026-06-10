package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/VelVit24/todo-api/database"
	"github.com/VelVit24/todo-api/models"
	"github.com/VelVit24/todo-api/service"
)

type Handler struct {
	DB *sql.DB
}

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
	if !service.ValidateLogin(user) {
		http.Error(w, "error", http.StatusBadRequest)
		return
	}
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
	if !service.ValidateLogin(user) {
		http.Error(w, "error", http.StatusBadRequest)
		return
	}
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
	id_user := r.Context().Value(service.UserIDKey).(int)
	switch r.Method {
	case "POST":
		h.HandleInsertTodos(w, r, id_user)
	case "GET":
		h.HandleGetTodos(w, r, id_user)
	}
}
func (h *Handler) HandleInsertTodos(w http.ResponseWriter, r *http.Request, id int) {
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

	err = database.InsertTodo(h.DB, id, &todo)
	if err != nil {
		log.Println(err)
		return
	}
	service.WriteJson(w, 200, todo)
}
func (h *Handler) HandleGetTodos(w http.ResponseWriter, r *http.Request, id int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	todos, err := database.GetTodos(h.DB, id, page, limit)
	if err != nil {
		log.Println(err)
		return
	}
	total, err := database.CountTodos(h.DB, id)
	if err != nil {
		log.Println(err)
		return
	}
	resp := models.TodosResponse{
		Data:  todos,
		Page:  page,
		Limit: limit,
		Total: total,
	}
	service.WriteJson(w, 200, resp)
}

func (h *Handler) HandleTodosInd(w http.ResponseWriter, r *http.Request) {
	id_user := r.Context().Value(service.UserIDKey).(int)
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/todos/"))
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	switch r.Method {
	case "DELETE":
		h.HandleDeleteTodos(w, r, id, id_user)
	case "PUT":
		h.HandleUpdateTodos(w, r, id, id_user)
	}

}
func (h *Handler) HandleDeleteTodos(w http.ResponseWriter, r *http.Request, id, id_user int) {
	err := database.DeleteTodo(h.DB, id, id_user)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "not found", 404)
		}
		log.Println(err)
		return
	}
	w.WriteHeader(204)
}
func (h *Handler) HandleUpdateTodos(w http.ResponseWriter, r *http.Request, id, id_user int) {
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
	err = database.UpdateTodo(h.DB, id, id_user, &todo)
	todo.Id = id
	if err != nil {
		log.Println(err)
		return
	}
	service.WriteJson(w, 200, todo)
}

package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/VelVit24/todo-api/database"
	"github.com/VelVit24/todo-api/models"
	"github.com/VelVit24/todo-api/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) PostLogin(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !service.ValidateLogin(user) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validate error"})
		return
	}
	ok, err := database.CheckUser(h.DB, &user)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect email or password"})
			return
		}
		log.Println(err)
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect email or password"})
		return
	}
	tokenString, err := service.GenToken(user.Id)
	if err != nil {
		log.Println(err)
		return
	}
	token := models.Token{Token: tokenString}
	c.JSON(200, token)
}

func (h *Handler) PostRegister(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !service.ValidateLogin(user) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validate error"})
		return
	}
	err := database.InsertUser(h.DB, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenString, err := service.GenToken(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := models.Token{Token: tokenString}
	c.JSON(200, token)
}

func (h *Handler) PostTodos(c *gin.Context) {
	todo := models.Todo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := c.Get(service.UserIDKey)
	err := database.InsertTodo(h.DB, id.(int), &todo)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, todo)
}

func (h *Handler) PutTodos(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := models.Todo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id_user, _ := c.Get(service.UserIDKey)
	err = database.UpdateTodo(h.DB, id, id_user.(int), &todo)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Status(404)
			return
		}
		log.Println(err)
		return
	}
	c.JSON(200, todo)
}

func (h *Handler) DeleteTodos(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id_user, _ := c.Get(service.UserIDKey)
	err = database.DeleteTodo(h.DB, id, id_user.(int))
	if err != nil {
		if err == sql.ErrNoRows {
			c.Status(404)
			return
		}
		log.Println(err)
		return
	}
	c.Status(204)
}

func (h *Handler) GetTodos(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	id_user, _ := c.Get(service.UserIDKey)
	todos, err := database.GetTodos(h.DB, id_user.(int), page, limit)
	if err != nil {
		log.Println(err)
		return
	}
	total, err := database.CountTodos(h.DB, id_user.(int))
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"data":  todos,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

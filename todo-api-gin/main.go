package main

import (
	"github.com/VelVit24/todo-api/database"
	"github.com/VelVit24/todo-api/handlers"
	"github.com/VelVit24/todo-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	db := database.ConnDB()
	defer db.Close()
	h := handlers.Handler{DB: db}
	router := gin.Default()
	router.Use(middleware.RateLimitMiddleware())
	router.POST("/login", h.PostLogin)
	router.POST("/register", h.PostRegister)
	router.POST("/todos", middleware.AuthMiddleware(), h.PostTodos)
	router.PUT("/todos/:id", middleware.AuthMiddleware(), h.PutTodos)
	router.DELETE("/todos/:id", middleware.AuthMiddleware(), h.DeleteTodos)
	router.GET("/todos", middleware.AuthMiddleware(), h.GetTodos)
	return router
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}

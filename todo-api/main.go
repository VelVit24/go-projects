package main

import (
	"log"
	"net/http"

	"github.com/VelVit24/todo-api/database"
	"github.com/VelVit24/todo-api/handlers"
)

func main() {
	db := database.ConnDB()
	defer db.Close()
	h := handlers.Handler{DB: db}
	http.HandleFunc("/register", h.HandleReg)
	http.HandleFunc("/login", h.HandleLogin)
	http.HandleFunc("/todos", h.HandleTodos)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

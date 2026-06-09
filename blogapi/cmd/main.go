package main

import (
	"log"
	"net/http"

	"github.com/VelVit24/blogapi/internal/database"
	"github.com/VelVit24/blogapi/internal/handlers"
)

func main() {
	db := database.ConnDB()
	defer db.Close()
	h := &handlers.Handler{DB: db}

	http.HandleFunc("/posts", h.HandlePosts)
	http.HandleFunc("/posts/", h.HandlePostsInd)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

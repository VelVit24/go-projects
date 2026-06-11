package main

import (
	"github.com/VelVit24/blogapi/internal/database"
	"github.com/VelVit24/blogapi/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnDB()
	defer db.Close()
	h := &handlers.Handler{DB: db}

	router := gin.Default()
	router.POST("/posts", h.PostPosts)
	router.GET("posts/:id", h.GetPostsId)
	router.PUT("/posts/:id", h.PutPostsId)
	router.DELETE("/posts/:id", h.DeletePostsId)
	router.GET("posts", h.GetPosts)

	router.Run(":8080")
}

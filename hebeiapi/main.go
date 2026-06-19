package main

import (
	"database/sql"

	"github.com/VelVit24/database"
	"github.com/VelVit24/handlers"
	"github.com/VelVit24/repository"
	"github.com/VelVit24/service"
	"github.com/gin-gonic/gin"
)

func App(db *sql.DB) *handlers.Handler {
	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := handlers.NewHandler(s)
	return h
}

func main() {
	db := database.DbConn()
	defer db.Close()
	h := App(db)
	r := gin.Default()
	blog := r.Group("/api/blog")
	{
		blog.GET("/categories", h.GetCategories)
		blog.GET("/posts", h.GetPosts)
		blog.GET("/posts/:slug", h.GetPostSlug)
		blog.GET("/posts/:slug/comments", h.GetComments)
		blog.POST("/posts/:slug/comments", h.CreateComment)
		blog.GET("/posts/:slug/cover", h.GetPostImage)
	}

	r.POST("/api/contact/send_mail", h.PostEmail)

	r.GET("/api/loaders/getall", h.GetLoaders)
	r.GET("/api/loaders/getimage/:loader_id", h.GetLoaderImage)
	r.GET("/api/manual_loaders/getall", h.GetManualLoaders)
	r.GET("/api/manual_loaders/getimage/:loader_id", h.GetManualLoaderImage)

	r.Run(":8080")
}

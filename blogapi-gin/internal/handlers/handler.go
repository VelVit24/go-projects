package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/VelVit24/blogapi/internal/database"
	"github.com/VelVit24/blogapi/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) PostPosts(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := database.InsertBlog(h.DB, &blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, blog)
}
func (h *Handler) GetPostsId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blog, err := database.GetBlog(h.DB, id)
	c.JSON(200, blog)
}

func (h *Handler) PutPostsId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var blog models.Blog
	err = c.ShouldBindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.UpdateBlog(h.DB, id, &blog)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, blog)
}

func (h *Handler) DeletePostsId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DeleteBlog(h.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(204)
}

func (h *Handler) GetPosts(c *gin.Context) {
	param := c.Query("term")
	blogs, err := database.GetAllBlogs(h.DB, param)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, blogs)
}

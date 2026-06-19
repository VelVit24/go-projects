package handlers

import (
	"strconv"

	"github.com/VelVit24/models"
	"github.com/VelVit24/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	serv *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{serv: serv}
}

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := h.serv.GetCategories()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, categories)
}

func (h *Handler) GetPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("perPage"))
	categories := c.QueryArray("category")
	posts, err := h.serv.GetPosts(page, perPage, categories)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, posts)
}

func (h *Handler) GetPostSlug(c *gin.Context) {
	slug := c.Param("slug")
	post, err := h.serv.GetPostSlug(slug)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, post)
}
func (h *Handler) GetComments(c *gin.Context) {

}
func (h *Handler) CreateComment(c *gin.Context) {
	//slug := c.Param("slug")
	request := models.CommentCreateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, "bad request")
	}
	responce, err := h.serv.CreateComment(&request)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(204, responce)
}
func (h *Handler) GetPostImage(c *gin.Context) {
	slug := c.Param("slug")

}

func (h *Handler) PostEmail(c *gin.Context) {
	request := models.ContactMail{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, "bad request")
	}
	status, err := h.serv.SendEmail(&request)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "email sent"})
}

func (h *Handler) GetLoaders(c *gin.Context) {
	loaders, err := h.serv.GetLoaders()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, loaders)
}

func (h *Handler) GetManualLoaders(c *gin.Context) {
	loaders, err := h.serv.GetManualLoaders()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, loaders)
}
func (h *Handler) GetLoaderImage(c *gin.Context) {

s
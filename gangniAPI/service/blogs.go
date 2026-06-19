package service

import (
	"github.com/VelVit24/models"
	"github.com/VelVit24/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetCategories() ([]models.BlogCategory, error) {
	categories, err := s.repo.SelectBlogCategories()
	return categories, err
}

func (s *Service) GetPosts(page, limit int, categories []string) (models.BlogPostsResponse, error) {
	posts, err := s.repo.SelectBlogPosts(page, limit, categories)
	return posts, err
}

func (s *Service) GetPostSlug(slug string) (models.BlogPostResponse, error) {
	return s.repo.SelectBlogPostSlug(slug)
}
func (s *Service) GetComments(slug string) ([]models.Comment, error) {
	return s.repo.SelectComments(slug)
}
func (s *Service) CreateComment(comment *models.CommentCreateRequest) (models.Comment, error) {
	return s.repo.InsertComment(comment)
}
func (s *Service) GetPostImage(slug string) (string, string, error) {
	return s.repo.SelectPostImage(slug)
}

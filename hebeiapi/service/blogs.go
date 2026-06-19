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
func (s *Service) GetComments(slug string) ([]*models.Comment, error) {
	comments, err := s.repo.SelectComments(slug)
	if err != nil {
		return nil, err
	}
	responce := []*models.Comment{}
	items := map[int]*models.Comment{}
	for i := range comments {
		comments[i].Replies = []*models.Comment{}
		items[comments[i].Id] = &comments[i]

	}
	for i := range comments {
		com := &comments[i]
		if com.ParentId == 0 {
			responce = append(responce, com)
			continue
		}
		parent := items[com.ParentId]
		if parent != nil {
			parent.Replies = append(parent.Replies, com)
		}
	}
	return responce, err
}

func (s *Service) CreateComment(slug string, comment *models.CommentCreateRequest) (models.Comment, error) {
	return s.repo.InsertComment(slug, comment)
}

func (s *Service) GetPostImage(slug string) (string, error) {
	linux, _, err := s.repo.SelectPostImage(slug)
	if err != nil {
		return "", err
	}
	return linux, nil

}

package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VelVit24/models"
	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SelectBlogCategories() ([]models.BlogCategory, error) {
	rows, err := r.db.Query("select id, name, slug from blog_category")
	if err != nil {
		return nil, err
	}
	categories := []models.BlogCategory{}
	for rows.Next() {
		cat := models.BlogCategory{}
		err := rows.Scan(&cat.Id, &cat.Name, &cat.Slug)
		if err != nil {
			log.Println(err.Error())
		}
		categories = append(categories, cat)
	}
	return categories, nil
}
func (r *Repository) SelectBlogPosts(page, limit int, categories []string) (models.BlogPostsResponse, error) {
	selectQuery := `select
	authors.avatar, authors.id, authors.name,
	blogPicturePathLinux, blogPicturePathWindows,
	categories.id, categories.name, categories.slug,
	excerpt, id, publishedAt, slug, title
	from posts left outer join categories on posts.category_slug=categories.slug
	left outer join authors on posts.authors=authors.id
	where 1=1
	`

	total := 0
	totalQuery := "select count(*) from posts"
	query := ""

	ind := 1
	args := []any{}
	for _, v := range categories {
		query += fmt.Sprintf(" and category_slug=$%d", ind)
		args = append(args, v)
		ind++
	}
	err := r.db.QueryRow(totalQuery+query, args...).Scan(&total)
	if err != nil {
		return models.BlogPostsResponse{}, err
	}

	query += fmt.Sprintf(" offset $%d limit $%d", ind, ind+1)
	args = append(args, page, limit)
	rows, err := r.db.Query(selectQuery+query, args...)
	if err != nil {
		return models.BlogPostsResponse{}, err
	}
	postResponce := models.BlogPostsResponse{}
	for rows.Next() {
		post := models.BlogPostListItem{}
		err := rows.Scan(
			&post.Author.Avatar, &post.Author.Id, &post.Author.Name,
			&post.BlogPicturePathLinux, &post.BlogPicturePathWindows,
			&post.Category.Id, &post.Category.Name, &post.Category.Name,
			&post.Excerpt, &post.PublishedAt, &post.Slug, &post.Title,
		)
		if err != nil {
			log.Println(err.Error())
		}
		postResponce.Data = append(postResponce.Data, post)

	}
	postResponce.Meta.Page = page
	postResponce.Meta.PerPage = limit
	postResponce.Meta.Total = total
	postResponce.Meta.TotalPages = total / limit
	return postResponce, nil
}

func (r *Repository) SelectBlogPostSlug(slug string) (models.BlogPostResponse, error) {
	query := `select
	authors.avatar, authors.id, authors.name,
	blogPicturePathLinux, blogPicturePathWindows,
	categories.id, categories.name, categories.slug,
	content, id, publishedAt, slug, tags, title
	from posts left outer join categories on posts.category_slug=categories.slug
	left outer join authors on posts.authors=authors.id
	where posts.slug=$1
	`
	row := r.db.QueryRow(query, slug)
	post := models.BlogPostDetail{}
	err := row.Scan(
		&post.Author.Avatar, &post.Author.Id, &post.Author.Name,
		&post.BlogPicturePathLinux, &post.BlogPicturePathWindows,
		&post.Category.Id, &post.Category.Name, &post.Category.Name,
		&post.Content, &post.PublishedAt, &post.Slug, pq.Array(&post.Tags), &post.Title,
	)
	if err != nil {
		return models.BlogPostResponse{}, err
	}
	return models.BlogPostResponse{Data: post}, nil
}

func (r *Repository) SelectComments(slug string) ([]models.Comment, error) {
	comments := []models.Comment{}
	rows, err := r.db.Query("select content, createdAt, id, name, parentId from comments where post_slug=$1", slug)
	if err != nil {
		return comments, nil
	}
	for rows.Next() {
		comment := models.Comment{}
		err := rows.Scan(&comment.Content, &comment.CreatedAt, &comment.Id, &comment.Name, &comment.ParentId)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *Repository) InsertComment(comment *models.CommentCreateRequest) (models.Comment, error) {
	comResponce := models.Comment{Content: comment.Content, Name: comment.Name, ParentId: comment.ParentId}
	err := r.db.QueryRow(`insert into comments(content, name, parentId) values
	($1, $2, $3) returning createdAt, id`, comment.Content, comment.Name, comment.ParentId).Scan(
		&comResponce.CreatedAt, &comResponce.Id,
	)
	return comResponce, err
}

func (r *Repository) SelectPostImage(slug string) (string, string, error) {
	var imageLinux, imageWindows string
	err := r.db.QueryRow("select blogPicturePathLinux, blogPicturePathWindows from posts where slug=$1", slug).Scan(&imageLinux, &imageWindows)
	return imageLinux, imageWindows, err
}

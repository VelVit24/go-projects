package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VelVit24/models"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SelectBlogCategories() ([]models.BlogCategory, error) {
	rows, err := r.db.Query("select id, name, slug from categories")
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
	blog_picture_path_linux, blog_picture_path_windows,
	categories.id, categories.name, categories.slug,
	excerpt, posts.id, published_at, posts.slug, posts.title
	from posts left outer join categories on posts.category_id=categories.id
	left outer join authors on posts.author_id=authors.id
	`

	total := 0
	totalQuery := "select count(*) from posts left outer join categories on posts.category_id=categories.id"
	query := ""

	if len(categories) > 0 {
		query += " where 1=0"
	}
	ind := 1
	args := []any{}
	for _, v := range categories {
		query += fmt.Sprintf(" or categories.slug=$%d", ind)
		args = append(args, v)
		ind++
	}
	err := r.db.QueryRow(totalQuery+query, args...).Scan(&total)
	if err != nil {
		return models.BlogPostsResponse{}, err
	}

	query += fmt.Sprintf(" offset $%d limit $%d", ind, ind+1)
	args = append(args, (page-1)*limit, limit)
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
			&post.Category.Id, &post.Category.Name, &post.Category.Slug,
			&post.Excerpt, &post.Id, &post.PublishedAt, &post.Slug, &post.Title,
		)
		if err != nil {
			log.Println(err.Error())
		}
		postResponce.Data = append(postResponce.Data, post)

	}
	postResponce.Meta.Page = page
	postResponce.Meta.PerPage = limit
	postResponce.Meta.Total = total
	postResponce.Meta.TotalPages = (total + limit - 1) / limit
	return postResponce, nil
}

func (r *Repository) SelectBlogPostSlug(slug string) (models.BlogPostResponse, error) {
	query := `select
	authors.avatar, authors.id, authors.name,
	blog_picture_path_linux, blog_picture_path_windows,
	categories.id, categories.name, categories.slug,
	content, posts.id, published_at, posts.slug, title
	from posts left outer join categories on posts.category_id=categories.id
	left outer join authors on posts.author_id=authors.id
	
	where posts.slug=$1
	`
	row := r.db.QueryRow(query, slug)
	post := models.BlogPostDetail{}
	err := row.Scan(
		&post.Author.Avatar, &post.Author.Id, &post.Author.Name,
		&post.BlogPicturePathLinux, &post.BlogPicturePathWindows,
		&post.Category.Id, &post.Category.Name, &post.Category.Slug,
		&post.Content, &post.Id, &post.PublishedAt, &post.Slug, &post.Title,
	)
	if err != nil {
		return models.BlogPostResponse{}, err
	}
	rows, err := r.db.Query(`select tags.id, tags.name, tags.slug
	from posts
	left outer join post_tags on posts.id=post_tags.post_id
	left outer join tags on post_tags.tag_id=tags.id
	where posts.id=$1`, post.Id)
	if err != nil {
		return models.BlogPostResponse{}, err
	}
	for rows.Next() {
		tag := models.BlogTag{}
		err := rows.Scan(&tag.Id, &tag.Name, &tag.Slug)
		if err != nil {
			log.Println(err)
		}
		post.Tags = append(post.Tags, tag)
	}
	return models.BlogPostResponse{Data: post}, nil
}

func (r *Repository) SelectComments(slug string) ([]models.Comment, error) {
	comments := []models.Comment{}
	rows, err := r.db.Query(`select c.content, c.created_at, c.id, c.name, c.parent_id 
	from comments c left outer join posts on c.post_id=posts.id
	where posts.slug=$1`, slug)
	if err != nil {
		return comments, err
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

func (r *Repository) InsertComment(slug string, comment *models.CommentCreateRequest) (models.Comment, error) {
	id_post := 0
	err := r.db.QueryRow("select id from posts where slug=$1", slug).Scan(&id_post)
	if err != nil {
		return models.Comment{}, err
	}
	comResponce := models.Comment{Content: comment.Content, Name: comment.Name, ParentId: comment.ParentId}
	err = r.db.QueryRow(`insert into comments(content, name, parent_id, post_id) values
	($1, $2, $3, $4) returning created_at, id`, comment.Content, comment.Name, comment.ParentId, id_post).Scan(
		&comResponce.CreatedAt, &comResponce.Id,
	)
	return comResponce, err
}

func (r *Repository) SelectPostImage(slug string) (string, string, error) {
	var imageLinux, imageWindows string
	err := r.db.QueryRow("select blog_picture_path_linux, blog_picture_path_windows from posts where slug=$1", slug).Scan(&imageLinux, &imageWindows)
	return imageLinux, imageWindows, err
}

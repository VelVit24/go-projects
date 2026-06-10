package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/VelVit24/blogapi/internal/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func ConnDB() *sql.DB {
	connStr := "user=postgres password=080907 dbname=blogdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения", err)
	}
	return db
}

func InsertBlog(db *sql.DB, blog *models.Blog) error {
	err := db.QueryRow("insert into Blogs (title, content, category, tags, createdAt, updatedAt) values ($1,$2,$3,$4,$5,$6) returning id",
		blog.Title, blog.Content, blog.Category, blog.Tags, time.Now(), time.Now()).Scan(&blog.Id)
	return err
}
func GetBlog(db *sql.DB, id int) (*models.Blog, error) {
	blog := &models.Blog{}
	row := db.QueryRow("select * from Blogs where id = $1", id)
	err := row.Scan(
		&blog.Id,
		&blog.Title,
		&blog.Content,
		&blog.Category,
		pq.Array(&blog.Tags),
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return blog, err
}

func DeleteBlog(db *sql.DB, id int) error {
	res, err := db.Exec("delete from Blogs where id = $1", id)
	if err != nil {
		return err
	}
	if s, _ := res.RowsAffected(); s == 0 {
		return sql.ErrNoRows
	}
	return err
}

func UpdateBlog(db *sql.DB, id int, blog *models.Blog) error {
	row := db.QueryRow("update Blogs set title = $1, content = $2, category = $3, tags = $4, updatedAt = $5 where id = $6 returning *", blog.Title, blog.Content, blog.Category, blog.Tags, time.Now(), id)
	err := row.Scan(
		&blog.Id,
		&blog.Title,
		&blog.Content,
		&blog.Category,
		pq.Array(&blog.Tags),
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		}
		return err
	}
	return err
}

func GetAllBlogs(db *sql.DB, param string) (*[]models.Blog, error) {
	blogs := []models.Blog{}
	var row *sql.Rows
	var err error
	if len(param) == 0 {
		row, err = db.Query("select * from Blogs")
	} else {
		row, err = db.Query(`select * from Blogs 
			where 
			title ILIKE '%'||$1||'%' 
			or content ILIKE '%'||$1||'%' 
			or category ILIKE '%'||$1||'%' 
			or EXISTs (
				select 1
				from unnest(tags) as tag
				where tag ILIKE '%'||$1||'%' );`, param)
	}
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		blog := models.Blog{}
		err := row.Scan(
			&blog.Id,
			&blog.Title,
			&blog.Content,
			&blog.Category,
			pq.Array(&blog.Tags),
			&blog.CreatedAt,
			&blog.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		blogs = append(blogs, blog)
	}
	return &blogs, err
}

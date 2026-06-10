package database

import (
	"database/sql"
	"log"

	"github.com/VelVit24/todo-api/models"
	"github.com/VelVit24/todo-api/service"
	_ "github.com/lib/pq"
)

func ConnDB() *sql.DB {
	connStr := "user=postgres password=080907 dbname=tododb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InsertUser(db *sql.DB, user *models.User) error {
	hash, err := service.HashPassword(user.Password)
	if err != nil {
		return err
	}
	return db.QueryRow("insert into Users(name, email, password) values ($1, $2, $3) returning id",
		user.Name, user.Email, hash).Scan(&user.Id)
}

func CheckUser(db *sql.DB, user *models.User) (bool, error) {
	var hash string
	err := db.QueryRow("select id, password from Users where email=$1", user.Email).Scan(&user.Id, &hash)
	if err != nil {
		return false, err
	}
	return service.CheckPassword(user.Password, hash), nil
}
func InsertTodo(db *sql.DB, id int, todo *models.Todo) error {
	err := db.QueryRow("insert into Todos (id_user, title, description) values ($1,$2,$3) returning id", id, todo.Title, todo.Description).Scan(&todo.Id)
	return err
}
func UpdateTodo(db *sql.DB, id, id_user int, todo *models.Todo) error {
	res, err := db.Exec("update Todos set title=$1, description=$2 where id=$3 and id_user=$4", todo.Title, todo.Description, id, id_user)
	if err != nil {
		return err
	}
	if s, _ := res.RowsAffected(); s == 0 {
		return sql.ErrNoRows
	}
	return err
}
func DeleteTodo(db *sql.DB, id, id_user int) error {
	res, err := db.Exec("delete from Todos where id=$1 and id_user=$2", id, id_user)
	if err != nil {
		return err
	}
	if s, _ := res.RowsAffected(); s == 0 {
		return sql.ErrNoRows
	}
	return err
}
func GetTodos(db *sql.DB, id_user, page, limit int) ([]models.Todo, error) {
	todos := []models.Todo{}
	row, err := db.Query("select id, title, description from Todos where id_user=$1 offset $2 limit $3 ", id_user, (page-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		todo := models.Todo{}
		err := row.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Description,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func CountTodos(db *sql.DB, id_user int) (int, error) {
	total := 0
	err := db.QueryRow("select count(*) from Todos where id_user=$1", id_user).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

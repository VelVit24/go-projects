package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DbConn() *sql.DB {
	godotenv.Load()
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

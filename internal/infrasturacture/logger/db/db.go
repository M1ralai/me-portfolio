package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "db/posts.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

package repository

import (
	"github.com/M1ralai/me-portfolio/internal/modules/post/domain"
	"github.com/jmoiron/sqlx"
)

type SqliteRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) domain.PostRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (r SqliteRepository) List() ([]domain.Post, error) {
	var posts []domain.Post
	query := `
		SELECT * FROM post;
	`
	if err := r.db.Select(&posts, query); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r SqliteRepository) GetById(id int) (*domain.Post, error) {
	var resp domain.Post
	query := `
		SELECT * FROM post WHERE id = ?;
	`
	if err := r.db.Select(resp, query, id); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r SqliteRepository) Create(p domain.Post) error {
	query := `
		INSERT INTO post (title, content, excerpt, date)
		VALUES (:title, :content, :excerpt, datetime('now'));
	`
	_, err := r.db.NamedExec(query, p)
	return err
}

func (r SqliteRepository) Delete(id int) error {
	query := `
		DELETE FROM post WHERE id = ?
	`
	_, err := r.db.Exec(query, id)
	return err
}

func (r SqliteRepository) Update(p domain.Post) error {
	query := `
		UPDATE post
		SET title=:title, content=:content, excerpt=:excerpt, date=datetime('now')
		WHERE id =:id
	`
	_, err := r.db.NamedExec(query, p)
	return err
}

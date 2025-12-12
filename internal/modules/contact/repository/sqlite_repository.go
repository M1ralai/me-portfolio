package repository

import (
	"github.com/M1ralai/me-portfolio/internal/modules/contact/domain"
	"github.com/jmoiron/sqlx"
)

type SqliteRepository struct {
	db *sqlx.DB
}

func NewContactRepository(db *sqlx.DB) domain.ContactRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (r SqliteRepository) List() ([]domain.Contact, error) {
	var contacts []domain.Contact
	query := `
		SELECT * FROM contact;
	`
	if err := r.db.Select(&contacts, query); err != nil {
		return nil, err
	}
	return contacts, nil
}

func (r SqliteRepository) Create(c domain.Contact) error {
	query := `
		INSERT INTO post (email, name, surname, company, message)
		VALUES (:email, :name, :surname, :company, :message);
	`
	_, err := r.db.NamedExec(query, c)
	return err
}

func (r SqliteRepository) Delete(id int) error {
	query := `
		DELETE FROM contact WHERE id = ?
	`
	_, err := r.db.Exec(query, id)
	return err
}

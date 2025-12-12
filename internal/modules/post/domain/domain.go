package domain

type Post struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"  validate:"required"`
	Content string `json:"content" db:"content"  validate:"required"`
	Excerpt string `json:"excerpt" db:"excerpt"  validate:"required"`
	Date    string `json:"date" db:"date"`
}

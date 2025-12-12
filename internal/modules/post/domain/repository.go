package domain

type PostRepository interface {
	List() ([]Post, error)
	GetById(id int) (*Post, error)
	Create(p Post) error
	Delete(id int) error
	Update(p Post) error
}

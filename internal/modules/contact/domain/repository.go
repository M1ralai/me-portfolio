package domain

type ContactRepository interface {
	List() ([]Contact, error)
	Create(c Contact) error
	Delete(id int) error
}

package service

import "github.com/M1ralai/me-portfolio/internal/modules/contact/domain"

type ContactService interface {
	List() ([]domain.Contact, error)
	Create(c domain.Contact) error
	Delete(id int) error
}

type contactService struct {
	repo domain.ContactRepository
}

func NewService(repo domain.ContactRepository) ContactService {
	return &contactService{
		repo: repo,
	}
}

func (s contactService) List() ([]domain.Contact, error) {
	//TODO authantication kontrolü
	return s.repo.List()
}

func (s contactService) Create(c domain.Contact) error {
	return s.Create(c)
}

func (s contactService) Delete(id int) error {
	//TODO authantication kontrolü
	return s.repo.Delete(id)
}

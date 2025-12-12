package service

import (
	"context"

	"github.com/M1ralai/me-portfolio/internal/modules/post/domain"
)

type PostService interface {
	List() ([]domain.Post, error)
	GetById(id int) (*domain.Post, error)
	Create(ctx context.Context, post domain.Post) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, post domain.Post) error
}

type postService struct {
	repo domain.PostRepository
}

func NewService(repo domain.PostRepository) PostService {
	return &postService{
		repo: repo,
	}
}

func (s postService) List() ([]domain.Post, error) {
	resp, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s postService) GetById(id int) (*domain.Post, error) {
	resp, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (s postService) Create(ctx context.Context, post domain.Post) error {
	//TODO ctx auth
	err := s.repo.Create(post)
	if err != nil {
		return err
	}
	return nil
}

func (s postService) Delete(ctx context.Context, id int) error {
	//TODO ctx auth
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s postService) Update(ctx context.Context, post domain.Post) error {
	//TODO ctx auth
	err := s.repo.Update(post)
	if err != nil {
		return err
	}
	return nil
}

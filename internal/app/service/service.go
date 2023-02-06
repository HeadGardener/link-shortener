package service

import "github.com/HeadGardener/link-shortener/internal/app/repository"

type Service struct {
	repos *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		repos: repos,
	}
}

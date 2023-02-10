package service

import (
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"github.com/HeadGardener/link-shortener/internal/app/repository"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GenerateToken(userInput models.UserInput) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Shortener interface {
	CreateLink(link models.InputLink, userID string) error
}

type Service struct {
	Authorization
	Shortener
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Shortener:     NewShortenerService(repos),
	}
}

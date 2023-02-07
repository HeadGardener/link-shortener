package service

import (
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"github.com/HeadGardener/link-shortener/internal/app/repository"
)

type Authorization interface {
	CreateUser(user models.User) (interface{}, error)
	GenerateToken(userInput models.UserInput) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}

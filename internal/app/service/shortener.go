package service

import (
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"github.com/HeadGardener/link-shortener/internal/app/repository"
	"github.com/HeadGardener/link-shortener/pkg/shortener"
	"github.com/google/uuid"
	"time"
)

type ShortenerService struct {
	repos *repository.Repository
}

func NewShortenerService(repos *repository.Repository) *ShortenerService {
	return &ShortenerService{repos: repos}
}

func (s *ShortenerService) CreateLink(inputLink models.InputLink, userID string) (models.Link, error) {
	linkID := uuid.New().ID()
	var shortURL string

	if inputLink.CustomURL == "" {
		shortURL = shortener.GetShortURL(linkID)
	} else {
		shortURL = inputLink.CustomURL
	}

	link := models.Link{
		URL:        inputLink.URL,
		ShortURL:   shortener.BaseURL + shortURL,
		Identifier: shortURL,
		ID:         linkID,
		UserID:     userID,
		CreatedAt:  time.Now(),
	}

	return link, s.repos.Shortener.CreateLink(link)
}

func (s *ShortenerService) Redirect(identifier string) (string, error) {
	link, err := s.repos.Shortener.GetLink(identifier)
	if err != nil {
		return "", err
	}

	return link.URL, nil
}

func (s *ShortenerService) GetAll(userID string) ([]models.Link, error) {
	return s.repos.Shortener.GetAllLinks(userID)
}

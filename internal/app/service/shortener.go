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

func (s *ShortenerService) CreateLink(inputLink models.InputLink, userID string) error {
	linkID := uuid.New().ID()
	var shortURL string

	if inputLink.CustomURL == "" {
		shortURL = shortener.GetShortURL(linkID)
	} else {
		shortURL = shortener.BaseURL + inputLink.CustomURL
	}

	link := models.Link{
		URL:       inputLink.URL,
		ShortURL:  shortURL,
		CustomURL: inputLink.CustomURL,
		ID:        linkID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	return s.repos.Shortener.CreateLink(link)
}

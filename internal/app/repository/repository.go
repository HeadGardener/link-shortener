package repository

import (
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUser(userInput models.UserInput) (string, error)
}

type Shortener interface {
	CreateLink(link models.Link) error
	GetLink(identifier string) (models.Link, error)
	GetAllLinks(userID string) ([]models.Link, error)
}

type Repository struct {
	Authorization
	Shortener
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
		Shortener:     NewShortenerMongo(db),
	}
}

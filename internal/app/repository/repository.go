package repository

import (
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user models.User) (interface{}, error)
	GetUser(userInput models.UserInput) (string, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
	}
}

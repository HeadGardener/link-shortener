package repository

import (
	"context"
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShortenerMongo struct {
	db *mongo.Database
}

func NewShortenerMongo(db *mongo.Database) *ShortenerMongo {
	return &ShortenerMongo{db: db}
}

func (r *ShortenerMongo) CreateLink(link models.Link) error {
	coll := r.db.Collection(LinksCollection)

	_, err := coll.InsertOne(context.TODO(), link)
	if err != nil {
		return err
	}

	return nil
}

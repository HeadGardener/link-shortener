package repository

import (
	"context"
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *ShortenerMongo) GetLink(identifier string) (models.Link, error) {
	coll := r.db.Collection(LinksCollection)

	filter := bson.D{{"identifier", identifier}}
	opts := options.FindOne()

	var link models.Link
	err := coll.FindOne(context.TODO(), filter, opts).Decode(&link)

	return link, err
}

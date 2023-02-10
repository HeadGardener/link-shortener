package repository

import (
	"context"
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthMongo struct {
	db *mongo.Database
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) CreateUser(user models.User) error {
	coll := r.db.Collection(UsersCollection)

	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthMongo) GetUser(userInput models.UserInput) (string, error) {
	coll := r.db.Collection(UsersCollection)

	filter := bson.D{{"username", userInput.Username}}
	opts := options.FindOne()

	var user models.User
	err := coll.FindOne(context.TODO(), filter, opts).Decode(&user)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

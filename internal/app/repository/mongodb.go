package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Config struct {
	Host   string
	Port   string
	DBName string
}

func NewMongoDB(config Config) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	db := client.Database(config.DBName)

	return db, nil
}

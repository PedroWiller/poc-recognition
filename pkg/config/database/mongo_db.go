package database

import (
	"context"
	"time"

	"poc-recognition/pkg/config"
	"poc-recognition/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	db *mongo.Client
}

func NewMongoDB() *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MONGO_URI))
	if err != nil {
		logger.Error("Error to connectiing in mongo", err.Error())
		return nil
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Error("Error to connectiing in mongo", err.Error())

		return nil
	}

	return &MongoDB{
		db: client,
	}
}

func (m *MongoDB) Collection() *mongo.Collection {
	return m.db.Database(config.MONGODB_NAME).Collection(config.MONGO_COLLECTION)
}

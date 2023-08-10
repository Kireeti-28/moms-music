package database

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getClient() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return &mongo.Client{}, err
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_URI")))
	if err != nil {
		return &mongo.Client{}, err
	}

	return client, nil
}

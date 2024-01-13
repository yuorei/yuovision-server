package db

import (
	"context"
	"log"

	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Database *mongo.Database
}

func NewMongoDB() *DB {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatalf("You must set your 'MONGODB_URI' environment variable.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("mongo.Connect: %s", err.Error())
	}

	database := client.Database("yuorei")
	if database == nil {
		log.Fatalf("database is nil")
	}
	return &DB{
		Database: database,
	}
}

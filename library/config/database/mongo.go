package database

import (
	"context"
	"fmt"
	"time"

	"github.com/Marvit-Solutions/csw-golang/library/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDBMongo(env config.Config) (*mongo.Database, error) {
	address := env.GetString("mongodb.address")
	port := env.GetString("mongodb.port")
	database := env.GetString("mongodb.database")

	uri := "mongodb://" + address + ":" + port

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Set a timeout for connection
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db := client.Database(database)

	return db, nil
}

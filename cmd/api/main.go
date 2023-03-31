package main

import (
	"context"
	"os"
	"time"

	"github.com/Ruthvik10/gosocial-user-mgmt/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal(err, nil)
	}
}

func main() {
	mongoURI := os.Getenv("MONGODB_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := dbConnect(ctx, mongoURI)
	if err != nil {
		logger.Fatal(err, nil)
	}
	err = client.Ping(ctx, nil) // ping to check if the mongoDB server is reachable and the client has established the connection successfully.
	if err != nil {
		logger.Fatal(err, nil)
	}
	logger.Print("Connected to the database", nil)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			logger.Fatal(err, nil)
		}
	}()
}

func dbConnect(ctx context.Context, URI string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI).SetReadPreference(readpref.Nearest())) // client reads from the nearest node by default which improves the read performance.
	if err != nil {
		return nil, err
	}
	return client, nil
}

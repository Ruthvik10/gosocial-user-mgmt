package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Ruthvik10/go-shared-library/logger"
	"github.com/Ruthvik10/gosocial-user-mgmt/internal/store"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type config struct {
	port string
	env  string
}
type application struct {
	cfg    config
	logger interface {
		Print(message string, properties map[string]any)
		Error(err error, properties map[string]any)
		Fatal(err error, properties map[string]any)
	}
	store store.Store
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg := config{
		port: os.Getenv("PORT"),
		env:  os.Getenv("ENV"),
	}
	app := application{
		cfg:    cfg,
		logger: logger.New(),
	}
	mongoURI := os.Getenv("MONGODB_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := dbConnect(ctx, mongoURI)
	if err != nil {
		app.logger.Fatal(err, nil)
	}
	err = client.Ping(ctx, nil) // ping to check if the mongoDB server is reachable and the client has established the connection successfully.
	if err != nil {
		app.logger.Fatal(err, nil)
	}
	app.logger.Print("Connected to the database", nil)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			app.logger.Fatal(err, nil)
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

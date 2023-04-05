package store

import (
	"context"
	"time"

	"github.com/Ruthvik10/gosocial-user-mgmt/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	client *mongo.Client
}

func (s *UserStore) Create(user *models.User) (*models.User, error) {
	coll := s.client.Database("userDB").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := coll.InsertOne(ctx, user)
	user.ID, _ = result.InsertedID.(primitive.ObjectID)

	if err != nil {
		return nil, err
	}
	return user, nil
}

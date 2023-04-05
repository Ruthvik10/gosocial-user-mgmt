package store

import (
	"github.com/Ruthvik10/gosocial-user-mgmt/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	User models.UserStore
}

func NewStore(client *mongo.Client) *Store {
	return &Store{
		User: &UserStore{client},
	}
}

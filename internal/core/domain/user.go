package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	Username       string               `bson:"username"`
	Email          string               `bson:"email"`
	Password       string               `bson:"password"`
	ProfilePicture *string              `bson:"profilePicture,omitempty"`
	Bio            *string              `bson:"bio, omitempty"`
	Birthdate      time.Time            `bson:"birthdate"`
	JoinedDate     time.Time            `bson:"joinedDate"`
	Followers      []primitive.ObjectID `bson:"followers,omitempty"`
	Following      []primitive.ObjectID `bson:"following,omitempty"`
}

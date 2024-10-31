package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sessions struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    primitive.ObjectID `bson:"user_id"`
	SessionID primitive.ObjectID `bson:"session_id"`
}

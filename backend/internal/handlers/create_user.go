package handlers

import (
	"context"

	"github.com/endingwithali/fitnessapp/backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(ctx context.Context, mongoDatabase *mongo.Database, discordUserID string) (primitive.ObjectID, error) {
	collection := mongoDatabase.Collection("users")
	userModel := models.User{
		ID:        primitive.NewObjectID(),
		DiscordID: discordUserID,
		UserID:    primitive.NewObjectID(),
	}
	_, err := collection.InsertOne(ctx, userModel)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return userModel.UserID, nil
}

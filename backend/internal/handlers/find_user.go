package handlers

import (
	"context"

	"github.com/endingwithali/fitnessapp/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ValidateUserInDatabase(ctx context.Context, mongoDatabase *mongo.Database, discordUserID string) (primitive.ObjectID, error) {
	collection := mongoDatabase.Collection("users")
	userObject := bson.D{{"discord_id", discordUserID}}

	var result models.Sessions
	err := collection.FindOne(ctx, userObject).Decode(&result)
	if err != nil {
		return primitive.NilObjectID, err
	}
	// if len(result) != 1 {
	// 	// [TODO] figure out how to create custom error type so when this file returns errors, I an do an error check on level up for specific error types
	// 	return fmt.Errorf("error: more than one user found with that discord id")
	// }
	return result.UserID, nil
}

// func FindUserByDiscordID(ctx context.Context, mongoDatabase *mongo.Database, discordUserID string)

func FindUser(ctx context.Context, userID string, mongoDatabase *mongo.Database) (models.User, error) {
	return models.User{}, nil
}

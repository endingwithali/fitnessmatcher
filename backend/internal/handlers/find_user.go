package handlers

import (
	"context"
	"fmt"

	"github.com/endingwithali/fitnessapp/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ValidateUserInDatabase(ctx context.Context, userID string, mongoDatabase *mongo.Database) error {
	collection := mongoDatabase.Collection("users")
	userObject := bson.D{{"discord_id", userID}}

	var result bson.M
	err := collection.FindOne(ctx, userObject).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return err
	}
	if err != nil {
		return err
	}
	if len(result) != 1 {
		return fmt.Errorf("error: more than one user found with that discord id")
	}

	return nil
}

func FindUser(ctx context.Context, userID string, mongoDatabase *mongo.Database) (models.User, error) {
	return models.User{}, nil
}

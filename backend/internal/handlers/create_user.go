package handlers

import (
	"context"
	"log"

	"github.com/endingwithali/fitnessapp/backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(ctx context.Context, mongoDatabase *mongo.Database, userModel models.User) error {
	collection := mongoDatabase.Collection("users")
	insertResult, err := collection.InsertOne(ctx, userModel)
	if err != nil {
		return err
	}
	log.Println(insertResult.InsertedID)
	return nil
}

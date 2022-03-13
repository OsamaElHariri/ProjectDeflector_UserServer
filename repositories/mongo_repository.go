package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"projectdeflector.users/users"
)

type MongoRepository struct {
	client *mongo.Client
	ctx    *context.Context
}

func (repo MongoRepository) InsertUser(uuid string) {
	repo.client.Database("user_management").Collection("users").InsertOne(*repo.ctx, bson.D{
		{Key: "uuid", Value: uuid},
	})
}

func (repo MongoRepository) FindUser(uuid string) (users.User, error) {
	var result users.User

	filter := bson.D{{Key: "uuid", Value: uuid}}
	repo.client.Database("user_management").Collection("users").CountDocuments(*repo.ctx, filter)
	err := repo.client.Database("user_management").Collection("users").FindOne(*repo.ctx, filter).Decode(&result)

	if err != nil {
		return result, err
	}
	return result, nil
}

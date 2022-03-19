package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	client *mongo.Client
	ctx    context.Context
}

type UserResult struct {
	Id        string `bson:"_id"`
	Uuid      string
	Nickname  string
	Color     string
	GameStats DbGameStat `bson:"game_stats"`
}

type UserInsertRequest struct {
	Uuid     string
	Nickname string
	Color    string
}

func (repo MongoRepository) InsertUser(user UserInsertRequest) {
	repo.client.Database("user_management").Collection("users").InsertOne(repo.ctx, bson.D{
		{Key: "uuid", Value: user.Uuid},
		{Key: "nickname", Value: user.Nickname},
		{Key: "color", Value: user.Color},
	})
}

func (repo MongoRepository) FindUser(uuid string) (UserResult, error) {
	var result UserResult

	filter := bson.D{{Key: "uuid", Value: uuid}}
	err := repo.client.Database("user_management").Collection("users").FindOne(repo.ctx, filter).Decode(&result)

	if err != nil {
		return result, err
	}
	return result, nil
}

func (repo MongoRepository) UpdateUser(user UserInsertRequest) (UserInsertRequest, error) {
	filter := bson.D{{Key: "uuid", Value: user.Uuid}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "uuid", Value: user.Uuid},
		{Key: "nickname", Value: user.Nickname},
		{Key: "color", Value: user.Color},
	}}}
	repo.client.Database("user_management").Collection("users").UpdateOne(repo.ctx, filter, update)

	return UserInsertRequest{
		Uuid:     user.Uuid,
		Nickname: user.Nickname,
		Color:    user.Color,
	}, nil
}

type DbGameStat struct {
	Games int
	Wins  int
}

func (repo MongoRepository) UpdateUserStats(id string, statUpdate DbGameStat) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "game_stats", Value: statUpdate},
	}}}
	_, err = repo.client.Database("user_management").Collection("users").UpdateOne(repo.ctx, filter, update)
	return err
}

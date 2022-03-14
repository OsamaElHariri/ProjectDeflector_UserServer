package repositories

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepositoryFactory struct {
	client *mongo.Client
}

func getMongoRepositoryFactory() MongoRepositoryFactory {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://some_user:password@127.0.0.1:27017"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return MongoRepositoryFactory{
		client: client,
	}
}

func (factory MongoRepositoryFactory) GetRepository() (Repository, func(), error) {
	client := factory.client
	ctx, cancelContext := context.WithTimeout(context.TODO(), 10*time.Second)

	cleanUpFunc := func() {
		cancelContext()
	}

	repo := MongoRepository{
		client: client,
		ctx:    ctx,
	}

	return repo, cleanUpFunc, nil
}

type MongoRepository struct {
	client *mongo.Client
	ctx    context.Context
}

func (repo MongoRepository) InsertUser(uuid string) {
	repo.client.Database("user_management").Collection("users").InsertOne(repo.ctx, bson.D{
		{Key: "uuid", Value: uuid},
	})
}

type FindUserResult struct {
	Uuid     string
	Nickname string
}

func (repo MongoRepository) FindUser(uuid string) (FindUserResult, error) {
	var result FindUserResult

	filter := bson.D{{Key: "uuid", Value: uuid}}
	err := repo.client.Database("user_management").Collection("users").FindOne(repo.ctx, filter).Decode(&result)

	if err != nil {
		return result, err
	}
	return result, nil
}

type UpdateUserResult struct {
	Uuid     string
	Nickname string
}

func (repo MongoRepository) UpdateUser(uuid string, nickname string) (UpdateUserResult, error) {
	filter := bson.D{{Key: "uuid", Value: uuid}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "uuid", Value: uuid},
		{Key: "nickname", Value: nickname},
	}}}
	repo.client.Database("user_management").Collection("users").UpdateOne(repo.ctx, filter, update)

	return UpdateUserResult{
		Uuid:     uuid,
		Nickname: nickname,
	}, nil
}

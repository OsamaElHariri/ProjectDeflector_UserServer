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

type DbUser struct {
	Uuid     string
	Nickname string
	Color    string
}

func (repo MongoRepository) InsertUser(user DbUser) {
	repo.client.Database("user_management").Collection("users").InsertOne(repo.ctx, bson.D{
		{Key: "uuid", Value: user.Uuid},
		{Key: "nickname", Value: user.Nickname},
		{Key: "color", Value: user.Color},
	})
}

func (repo MongoRepository) FindUser(uuid string) (DbUser, error) {
	var result DbUser

	filter := bson.D{{Key: "uuid", Value: uuid}}
	err := repo.client.Database("user_management").Collection("users").FindOne(repo.ctx, filter).Decode(&result)

	if err != nil {
		return result, err
	}
	return result, nil
}

func (repo MongoRepository) UpdateUser(user DbUser) (DbUser, error) {
	filter := bson.D{{Key: "uuid", Value: user.Uuid}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "uuid", Value: user.Uuid},
		{Key: "nickname", Value: user.Nickname},
		{Key: "color", Value: user.Color},
	}}}
	repo.client.Database("user_management").Collection("users").UpdateOne(repo.ctx, filter, update)

	return DbUser{
		Uuid:     user.Uuid,
		Nickname: user.Nickname,
		Color:    user.Color,
	}, nil
}

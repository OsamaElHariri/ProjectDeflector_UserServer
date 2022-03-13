package repositories

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	InsertUser(uuid string)
	FindUser(uuid string) (FindUserResult, error)
}

func GetRepository() (repository Repository, cleanUp func()) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://some_user:password@127.0.0.1:27017"))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancelContext := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	cleanUpFunc := func() {
		cancelContext()
		client.Disconnect(ctx)
	}

	repo := MongoRepository{
		client: client,
		ctx:    &ctx,
	}

	return repo, cleanUpFunc
}

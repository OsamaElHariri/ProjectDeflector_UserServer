package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://some_user:password@127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancelContext := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancelContext()
	defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	app.Post("/color", func(c *fiber.Ctx) error {
		payload := struct {
			PlayerId string `json:"playerId"`
			Color    string `json:"color"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"color": "ok",
		})
	})

	log.Fatal(app.Listen(":3000"))
}

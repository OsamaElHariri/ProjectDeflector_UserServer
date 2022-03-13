package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"projectdeflector.users/repositories"
	"projectdeflector.users/users"
)

func main() {
	app := fiber.New()

	repo, cleanup := repositories.GetRepository()
	defer cleanup()
	useCase := users.UseCase{
		Repo: repo,
	}

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		payload := struct {
			Uuid string `json:"uuid"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(400)
		}

		user, err := useCase.CreateNewAnonymousUser()
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"user": user,
		})
	})

	app.Get("/user/:uuid", func(c *fiber.Ctx) error {
		uuid := c.Params("uuid")

		user, err := useCase.GetUser(uuid)

		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"user": user,
		})
	})

	log.Fatal(app.Listen(":3005"))
}

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"projectdeflector.users/repositories"
	"projectdeflector.users/users"
)

func main() {
	app := fiber.New()

	repoFactory := repositories.GetRepositoryFactory()

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		repo, cleanup, err := repoFactory.GetRepository()
		defer cleanup()
		if err != nil {
			return c.SendStatus(400)
		}
		useCase := users.UseCase{
			Repo: repo,
		}

		user, err := useCase.CreateNewAnonymousUser()
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"user": user,
		})
	})

	app.Put("/user/:uuid", func(c *fiber.Ctx) error {
		uuid := c.Params("uuid")

		payload := struct {
			Nickname string `json:"nickname"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(400)
		}

		repo, cleanup, err := repoFactory.GetRepository()
		defer cleanup()
		if err != nil {
			return c.SendStatus(400)
		}
		useCase := users.UseCase{
			Repo: repo,
		}

		user, err := useCase.UpdateUser(uuid, payload.Nickname)
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"user": user,
		})
	})

	app.Get("/user/:uuid", func(c *fiber.Ctx) error {
		uuid := c.Params("uuid")

		repo, cleanup, err := repoFactory.GetRepository()
		defer cleanup()
		if err != nil {
			return c.SendStatus(400)
		}
		useCase := users.UseCase{
			Repo: repo,
		}

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

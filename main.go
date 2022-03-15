package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"projectdeflector.users/repositories"
	"projectdeflector.users/users"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	app := fiber.New()

	repoFactory := repositories.GetRepositoryFactory()

	app.Use("/", func(c *fiber.Ctx) error {
		repo, cleanup, err := repoFactory.GetRepository()
		if err != nil {
			return c.SendStatus(400)
		}

		defer cleanup()
		c.Locals("repo", repo)

		return c.Next()
	})

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		repo := c.Locals("repo").(repositories.Repository)
		useCase := users.UseCase{
			Repo: repo,
		}

		user, err := useCase.CreateNewAnonymousUser()
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"user": parseUser(user),
		})
	})

	app.Put("/user/:uuid", func(c *fiber.Ctx) error {
		repo := c.Locals("repo").(repositories.Repository)
		uuid := c.Params("uuid")

		payload := struct {
			Nickname string `json:"nickname"`
			Color    string `json:"color"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(400)
		}

		useCase := users.UseCase{
			Repo: repo,
		}

		user, err := useCase.UpdateUser(uuid, payload.Nickname, payload.Color)
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"user": parseUser(user),
		})
	})

	app.Get("/user/:uuid", func(c *fiber.Ctx) error {
		repo := c.Locals("repo").(repositories.Repository)
		uuid := c.Params("uuid")

		useCase := users.UseCase{
			Repo: repo,
		}

		user, err := useCase.GetUser(uuid)

		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"user": parseUser(user),
		})
	})

	log.Fatal(app.Listen(":3005"))
}

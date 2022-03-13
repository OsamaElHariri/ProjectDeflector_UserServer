package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"projectdeflector.users/repositories"
)

func main() {
	app := fiber.New()

	repo, cleanup := repositories.GetRepository()
	defer cleanup()

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

		repo.InsertUser(payload.Uuid)

		return c.JSON(fiber.Map{
			"uuid": payload.Uuid,
		})
	})

	app.Get("/user/:uuid", func(c *fiber.Ctx) error {
		uuid := c.Params("uuid")

		user, err := repo.FindUser(uuid)

		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"uuid":     user.Uuid,
			"nickanme": user.Nickname,
		})
	})

	log.Fatal(app.Listen(":3005"))
}

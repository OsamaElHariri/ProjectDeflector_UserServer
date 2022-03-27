package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"projectdeflector.users/repositories"
	"projectdeflector.users/users"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	err := godotenv.Load("env/." + env + ".env")
	if err != nil {
		log.Fatalf("could not load env vars")
	}

	rand.Seed(time.Now().UnixNano())
	app := fiber.New()
	app.Use(recover.New())

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

	app.Use("/", func(c *fiber.Ctx) error {
		userId := c.Get("x-user-id")
		if userId != "" {
			c.Locals("userId", userId)
		}
		return c.Next()
	})

	app.Get("/internal/auth/check", func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth != "" {
			token := strings.Replace(auth, "Bearer ", "", 1)
			secretInternalToken := os.Getenv("INTERNAL_TOKEN")
			if token != secretInternalToken {
				return c.SendStatus(403)
			}
		} else {
			return c.SendStatus(403)
		}
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Get("/auth/check", func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth != "" {
			token := strings.Replace(auth, "Bearer ", "", 1)
			userId, err := users.UseCase{}.ValidateAccessToken(token)
			if err != nil {
				return c.SendStatus(403)
			}
			c.Response().Header.Add("x-user-id", userId)
		} else {
			return c.SendStatus(403)
		}
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/internal/stats/games", func(c *fiber.Ctx) error {
		payload := struct {
			Updates []struct {
				PlayerId string `json:"playerId"`
				Games    int    `json:"games"`
				Wins     int    `json:"wins"`
			} `json:"updates"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(400)
		}

		repo := c.Locals("repo").(repositories.Repository)
		useCase := users.UseCase{
			Repo: repo,
		}

		statUpdates := []users.GameStatUpdate{}
		for i := 0; i < len(payload.Updates); i++ {
			statUpdates = append(statUpdates, users.GameStatUpdate{
				PlayerId: payload.Updates[i].PlayerId,
				Games:    payload.Updates[i].Games,
				Wins:     payload.Updates[i].Wins,
			})
		}

		useCase.UpdateUserStats(statUpdates)

		return c.JSON(fiber.Map{
			"processing": true,
		})
	})

	app.Post("/public/access", func(c *fiber.Ctx) error {
		payload := struct {
			Uuid string `json:"uuid"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(400)
		}
		repo := c.Locals("repo").(repositories.Repository)
		useCase := users.UseCase{
			Repo: repo,
		}

		token, err := useCase.GetAccessToken(payload.Uuid)
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"token": token,
		})
	})

	app.Post("/public/user", func(c *fiber.Ctx) error {
		repo := c.Locals("repo").(repositories.Repository)
		useCase := users.UseCase{
			Repo: repo,
		}

		uuid, user, err := useCase.CreateNewAnonymousUser()
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"uuid": uuid,
			"user": parseUser(user),
		})
	})

	app.Put("/user", func(c *fiber.Ctx) error {
		playerId := c.Locals("userId").(string)

		payload := struct {
			Nickname string `json:"nickname"`
			Color    string `json:"color"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(400)
		}

		repo := c.Locals("repo").(repositories.Repository)
		useCase := users.UseCase{
			Repo: repo,
		}

		user, err := useCase.UpdateUser(playerId, payload.Nickname, payload.Color)
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"user": parseUser(user),
		})
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		repo := c.Locals("repo").(repositories.Repository)
		useCase := users.UseCase{
			Repo: repo,
		}

		user, err := useCase.GetUser(id)

		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"user": parseUser(user),
		})
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		playerId := c.Locals("userId").(string)

		repo := c.Locals("repo").(repositories.Repository)
		useCase := users.UseCase{
			Repo: repo,
		}

		user, err := useCase.GetUser(playerId)

		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"user": parseUser(user),
		})
	})

	app.Get("/colors", func(c *fiber.Ctx) error {
		playerId := c.Locals("userId").(string)

		repo := c.Locals("repo").(repositories.Repository)
		useCase := users.UseCase{
			Repo: repo,
		}
		colors, err := useCase.GetUserColors(playerId)

		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(fiber.Map{
			"colors": colors,
		})
	})

	log.Fatal(app.Listen(":3006"))
}

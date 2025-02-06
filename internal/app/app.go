package app

import (
	"todo/config"

	"github.com/gofiber/fiber/v2"
)

func NewFiberApp(cfg config.Config) *fiber.App {
	app := fiber.New()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Todo App!")
	})

	return app
}

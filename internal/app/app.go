package app

import (
	"todo/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewFiberApp(cfg config.Config, db *gorm.DB) *fiber.App {
	app := fiber.New()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Todo App")
	})

	return app
}

package app

import (
	"todo/routes"

	"github.com/gofiber/fiber/v2"
)

func NewFiberApp() *fiber.App {
	app := fiber.New()

	// Initialize routes
	routes.SetupRoutes(app)

	return app
}

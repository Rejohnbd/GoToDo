package routes

import (
	"todo/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Todo App")
	})

	handlerTodo := handlers.NewTodoHandler()

	api := app.Group("/api")
	api.Get("/todos", handlerTodo.GetTodos)
	api.Post("/todos", handlerTodo.CreateTodo)
	api.Delete("/todos/:id", handlerTodo.DeleteTodo)
}

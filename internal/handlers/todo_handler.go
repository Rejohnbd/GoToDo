package handlers

import (
	"fmt"
	"todo/internal/models"
	"todo/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	Repo *repositories.TodoRepository
}

type TodoResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{Repo: repositories.NewTodoRepository()}
}

func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	todos := h.Repo.GetTodos()

	response := TodoResponse{
		Status:  true,
		Message: "Todo List",
		Data:    todos,
	}

	return c.JSON(response)
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	h.Repo.CreateTodo(&todo)
	return c.Status(201).JSON(todo)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	fmt.Println(id)
	h.Repo.DeleteTodo(uint(id))
	return c.SendStatus(204)
}

package handlers

import (
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

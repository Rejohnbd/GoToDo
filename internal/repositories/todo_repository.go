package repositories

import (
	"todo/config"
	"todo/internal/models"
)

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) GetTodos() []models.Todo {
	var todos []models.Todo
	config.DB.Find(&todos)
	return todos
}

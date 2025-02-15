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

func (r *TodoRepository) CreateTodo(todo *models.Todo) {
	config.DB.Create(todo)
}

func (r *TodoRepository) DeleteTodo(id uint) {
	config.DB.Delete(&models.Todo{}, id)
}

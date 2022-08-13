package ports

import "github.com/ismael3s/go-todo/internal/core/domain"

type TodoService interface {
	GetTodos() ([]*domain.Todo, error)
	GetTodo(id string) (*domain.Todo, error)
	Save(todo *domain.Todo) error
	Remove(id string) error
}

type TodoRepository interface {
	GetTodos() ([]*domain.Todo, error)
	GetTodo(id string) (*domain.Todo, error)
	Save(todo *domain.Todo) error
	Remove(id string) error
}

package ports

import "github.com/ismael3s/go-todo/internal/core/domain"

type TodoService interface {
	GetTodos() ([]*domain.Todo, error)
}

type TodoRepository interface {
	GetTodos() ([]*domain.Todo, error)
}

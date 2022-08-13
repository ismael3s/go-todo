package todo

import "github.com/ismael3s/go-todo/internal/core/domain"

type InMemoryTodoRepository struct {
	todos []*domain.Todo
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: []*domain.Todo{},
	}
}

func (r *InMemoryTodoRepository) GetTodos() ([]*domain.Todo, error) {
	return r.todos, nil
}

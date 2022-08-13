package todo

import "github.com/ismael3s/go-todo/internal/core/domain"

type InMemoryTodoRepository struct {
	todos []*domain.Todo
	err   error
}

func NewInMemoryTodoRepository(err error) *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: []*domain.Todo{},
		err:   err,
	}
}

func (r *InMemoryTodoRepository) GetTodos() ([]*domain.Todo, error) {
	if r.err != nil {
		return nil, r.err
	}

	return r.todos, nil
}

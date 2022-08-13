package todo

import (
	"errors"

	"github.com/ismael3s/go-todo/internal/core/domain"
)

type InMemoryTodoRepository struct {
	todos []*domain.Todo
	err   error
}

func NewInMemoryTodoRepository(todos []*domain.Todo, err error) *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: todos,
		err:   err,
	}
}

func (r *InMemoryTodoRepository) GetTodos() ([]*domain.Todo, error) {
	if r.err != nil {
		return nil, r.err
	}

	return r.todos, nil
}

func (r *InMemoryTodoRepository) GetTodo(id string) (*domain.Todo, error) {
	if r.err != nil {
		return nil, r.err
	}

	for _, todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return nil, errors.New("todo not found")
}

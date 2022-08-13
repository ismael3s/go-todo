package todo

import (
	"errors"
	"strings"

	"github.com/ismael3s/go-todo/internal/core/domain"
	"github.com/ismael3s/go-todo/internal/infra/utils"
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

func (r *InMemoryTodoRepository) Save(todo *domain.Todo) error {
	if r.err != nil {
		return r.err
	}

	if strings.TrimSpace(todo.ID) == "" {
		return r.create(todo)
	}

	return r.update(todo)
}

func (r *InMemoryTodoRepository) create(todo *domain.Todo) error {
	if r.err != nil {
		return r.err
	}

	todo.ID = utils.NewUUID()

	r.todos = append(r.todos, todo)

	return nil
}

func (r *InMemoryTodoRepository) update(todo *domain.Todo) error {
	if r.err != nil {
		return r.err
	}

	for i, t := range r.todos {
		if t.ID == todo.ID {
			r.todos[i] = todo
			return nil
		}
	}

	return errors.New("todo not found")
}

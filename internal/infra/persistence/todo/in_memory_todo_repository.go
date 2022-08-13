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

func (r *InMemoryTodoRepository) Remove(id string) error {
	if r.err != nil {
		return r.err
	}

	var index = -1

	for i, t := range r.todos {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("todo not found")
	}

	r.todos = r.removeIndex(index)

	return nil
}

func (r *InMemoryTodoRepository) removeIndex(index int) []*domain.Todo {
	return append(r.todos[:index], r.todos[index+1:]...)
}

func (r *InMemoryTodoRepository) Save(todo *domain.Todo) (*domain.Todo, error) {
	if r.err != nil {
		return nil, r.err
	}

	if strings.TrimSpace(todo.ID) == "" {
		return r.create(todo)
	}

	return r.update(todo)
}

func (r *InMemoryTodoRepository) create(todo *domain.Todo) (*domain.Todo, error) {
	if r.err != nil {
		return nil, r.err
	}

	todo.ID = utils.NewUUID()

	r.todos = append(r.todos, todo)

	return todo, nil
}

func (r *InMemoryTodoRepository) update(todo *domain.Todo) (*domain.Todo, error) {
	if r.err != nil {
		return nil, r.err
	}

	for i, t := range r.todos {
		if t.ID == todo.ID {
			r.todos[i] = todo
			return todo, nil
		}
	}

	return nil, errors.New("todo not found")
}

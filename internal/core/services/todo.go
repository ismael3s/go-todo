package services

import (
	"github.com/ismael3s/go-todo/internal/core/domain"
	"github.com/ismael3s/go-todo/internal/core/ports"
)

// s *TodoService ports.TodoService

type TodoService struct {
	todoRepository ports.TodoRepository
}

func (s *TodoService) GetTodos() ([]*domain.Todo, error) {
	return s.todoRepository.GetTodos()
}

func (s *TodoService) GetTodo(id string) (*domain.Todo, error) {
	todo, err := s.todoRepository.GetTodo(id)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) Save(todo *domain.Todo) error {
	return s.todoRepository.Save(todo)
}

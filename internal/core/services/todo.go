package services

import (
	"github.com/ismael3s/go-todo/internal/core/domain"
	"github.com/ismael3s/go-todo/internal/core/ports"
)

type TodoService struct {
	todoRepository ports.TodoRepository
}

func (s *TodoService) GetTodos() ([]*domain.Todo, error) {
	return s.todoRepository.GetTodos()
}

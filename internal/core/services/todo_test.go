package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ismael3s/go-todo/internal/core/domain"
	"github.com/ismael3s/go-todo/internal/core/ports"
	persistence "github.com/ismael3s/go-todo/internal/infra/persistence/todo"
)

func TestTodoService_GetTodos(t *testing.T) {
	type fields struct {
		todoRepository ports.TodoRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*domain.Todo
		wantErr bool
	}{
		{
			name: "should be able to return all todos",
			fields: fields{
				todoRepository: persistence.NewInMemoryTodoRepository([]*domain.Todo{}, nil),
			},
			want:    []*domain.Todo{},
			wantErr: false,
		},
		{
			name: "should not be able to return all todos",
			fields: fields{
				todoRepository: persistence.NewInMemoryTodoRepository([]*domain.Todo{}, errors.New("Failed to find todos")),
			},
			want:    []*domain.Todo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
			}
			got, err := s.GetTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoService.GetTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (err != nil) == tt.wantErr {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoService.GetTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoService_GetTodo(t *testing.T) {
	todos := []*domain.Todo{
		{
			ID:    "1",
			Title: "Todo 1",
		},
	}
	type fields struct {
		todoRepository ports.TodoRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Todo
		wantErr bool
	}{
		{
			name: "given an existing id should be able to return the todo",
			fields: fields{
				todoRepository: persistence.NewInMemoryTodoRepository(todos, nil),
			},
			args: args{
				id: "1",
			},
			want:    &domain.Todo{ID: "1", Title: "Todo 1"},
			wantErr: false,
		},
		{
			name: "given an unexisting id should throw an error",
			fields: fields{
				todoRepository: persistence.NewInMemoryTodoRepository(todos, nil),
			},
			args: args{
				id: "2",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
			}
			got, err := s.GetTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoService.GetTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoService.GetTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoService_Save(t *testing.T) {
	todosUpdate := []*domain.Todo{
		{
			ID:    "2",
			Title: "Todo 1",
		},
	}

	type fields struct {
		todoRepository ports.TodoRepository
	}
	type args struct {
		todo *domain.Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should be able to create a todo",
			fields: fields{
				todoRepository: persistence.NewInMemoryTodoRepository([]*domain.Todo{}, nil),
			},
			args: args{
				todo: &domain.Todo{
					Title: "Todo 1",
				},
			},
			wantErr: false,
		},
		{
			name: "should be able to update a todo",
			fields: fields{
				todoRepository: persistence.NewInMemoryTodoRepository(todosUpdate, nil),
			},
			args: args{
				todo: &domain.Todo{
					ID:    "2",
					Title: "Todo 2",
				},
			},
			wantErr: false,
		},
		{
			name: "should be return an error when given ID has no todo",
			fields: fields{
				todoRepository: persistence.NewInMemoryTodoRepository(todosUpdate, nil),
			},
			args: args{
				todo: &domain.Todo{
					ID:    "3",
					Title: "Todo 2",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
			}
			if err := s.Save(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("TodoService.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoService_Remove(t *testing.T) {
	todos := []*domain.Todo{
		{
			ID:    "1",
			Title: "Todo 1",
		},
	}

	type fields struct {
		todoRepository ports.TodoRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should be able to remove a todo",
			fields: fields{
				todoRepository: persistence.NewInMemoryTodoRepository(todos, nil),
			},
			args: args{
				id: "1",
			},
			wantErr: false,
		},
		{
			name: "should not be able to remove a todo if the id is not found",
			fields: fields{
				todoRepository: persistence.NewInMemoryTodoRepository(todos, nil),
			},
			args: args{
				id: "2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
			}
			if err := s.Remove(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoService.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

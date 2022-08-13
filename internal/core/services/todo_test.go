package services

import (
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
				todoRepository: persistence.NewInMemoryTodoRepository(),
			},
			want:    []*domain.Todo{},
			wantErr: false,
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoService.GetTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

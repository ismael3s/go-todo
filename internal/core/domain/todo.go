package domain

type Todo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewTodo(id string, name string) *Todo {
	return &Todo{
		ID:   id,
		Name: name,
	}
}

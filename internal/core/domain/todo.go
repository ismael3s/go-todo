package domain

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func NewTodo(id string, title string) *Todo {
	return &Todo{
		ID:    id,
		Title: title,
	}
}

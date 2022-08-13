package domain

import (
	"github.com/ismael3s/go-todo/internal/infra/utils"
)

type Todo struct {
	Base
	Title string `json:"title"`
}

func NewTodo(title string) *Todo {
	return &Todo{
		Base: Base{
			ID:        utils.NewUUID(),
			CreatedAt: utils.NowTime(),
			UpdatedAt: utils.NowTime(),
		},
		Title: title,
	}
}

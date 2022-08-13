package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
}

func NewTodoController() *TodoController {
	return &TodoController{}
}

func (ctrl *TodoController) GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

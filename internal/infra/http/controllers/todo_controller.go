package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ismael3s/go-todo/internal/core/domain"
	"github.com/ismael3s/go-todo/internal/core/ports"
)

type TodoController struct {
	service ports.TodoService
}

func NewTodoController(service ports.TodoService) *TodoController {
	return &TodoController{
		service: service,
	}
}

func (ctrl *TodoController) GetTodos(c *gin.Context) {
	todos, err := ctrl.service.GetTodos()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func (ctrl *TodoController) GetTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := ctrl.service.GetTodo(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (ctrl *TodoController) Save(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ctrl.service.Save(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

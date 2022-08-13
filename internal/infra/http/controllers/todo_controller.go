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
	var createTodo domain.Todo
	if err := c.ShouldBindJSON(&createTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	todo, err := ctrl.service.Save(&createTodo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func (ctrl *TodoController) Update(c *gin.Context) {
	var updateTodo domain.Todo
	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var id = c.Param("id")

	if updateTodo.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not match",
		})
		return
	}

	todo, err := ctrl.service.Update(&updateTodo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

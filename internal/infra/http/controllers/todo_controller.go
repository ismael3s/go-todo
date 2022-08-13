package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	teste string
}

func NewTodoController() *TodoController {
	return &TodoController{
		teste: "teste",
	}
}

func (ctrl *TodoController) GetTodos(c *gin.Context) {
	log.Println("GetTodos")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

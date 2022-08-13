package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ismael3s/go-todo/internal/infra/http/controllers"
)

type todoRouter struct {
	framework  *gin.Engine
	controller *controllers.TodoController
}

func NewTodoRouter(framework *gin.Engine, controller *controllers.TodoController) *todoRouter {
	return &todoRouter{
		framework:  framework,
		controller: controller,
	}
}

func (r *todoRouter) RegisterRoutes() {
	r.getTodos()
}

func (r *todoRouter) getTodos() {
	r.framework.GET("/todos", r.controller.GetTodos)
}

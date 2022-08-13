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
	r.getTodo()
	r.save()
	r.update()
}

func (r *todoRouter) getTodos() {
	r.framework.GET("/todos", r.controller.GetTodos)
}

func (r *todoRouter) getTodo() {
	r.framework.GET("/todos/:id", r.controller.GetTodo)
}

func (r *todoRouter) save() {
	r.framework.POST("/todos", r.controller.Save)
}

func (r *todoRouter) update() {
	r.framework.PUT("/todos/:id", r.controller.Update)
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ismael3s/go-todo/internal/infra/http/controllers"
	"github.com/ismael3s/go-todo/internal/infra/http/routes"
)

type RoutesHandler struct {
	framework *gin.Engine
}

func NewRoutesHandler(framework *gin.Engine) *RoutesHandler {
	return &RoutesHandler{framework}
}

func (r *RoutesHandler) RegisterRoutes() {
	todoController := controllers.NewTodoController()

	routes.NewTodoRouter(r.framework, todoController).RegisterRoutes()
}

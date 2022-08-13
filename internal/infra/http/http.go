package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ismael3s/go-todo/internal/core/domain"
	"github.com/ismael3s/go-todo/internal/core/services"
	"github.com/ismael3s/go-todo/internal/infra/http/controllers"
	"github.com/ismael3s/go-todo/internal/infra/http/routes"
	persistence "github.com/ismael3s/go-todo/internal/infra/persistence/todo"
)

type RoutesHandler struct {
	framework *gin.Engine
}

func NewRoutesHandler(framework *gin.Engine) *RoutesHandler {
	return &RoutesHandler{framework}
}

func (r *RoutesHandler) RegisterRoutes() {
	todoRepository := persistence.NewInMemoryTodoRepository([]*domain.Todo{}, nil)
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	routes.NewTodoRouter(r.framework, todoController).RegisterRoutes()
}

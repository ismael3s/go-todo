package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ismael3s/go-todo/internal/core/services"
	"github.com/ismael3s/go-todo/internal/infra/http/controllers"
	"github.com/ismael3s/go-todo/internal/infra/http/routes"
	"github.com/ismael3s/go-todo/internal/infra/persistence"
	repositories "github.com/ismael3s/go-todo/internal/infra/persistence/todo"
)

type RoutesHandler struct {
	framework *gin.Engine
}

func NewRoutesHandler(framework *gin.Engine) *RoutesHandler {
	return &RoutesHandler{framework}
}

func (r *RoutesHandler) RegisterRoutes() {
	db := persistence.GetPostgresConnection()
	todoRepository := repositories.NewPostgresTodoRepository(db)
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	routes.NewTodoRouter(r.framework, todoController).RegisterRoutes()
}

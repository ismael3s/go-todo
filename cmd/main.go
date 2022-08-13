package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ismael3s/go-todo/internal/infra/http/controllers"
	"github.com/ismael3s/go-todo/internal/infra/http/routes"
)

func main() {
	r := gin.Default()
	//
	todoController := controllers.NewTodoController()

	routes.NewTodoRouter(r, todoController).RegisterRoutes()
	r.Run()
}

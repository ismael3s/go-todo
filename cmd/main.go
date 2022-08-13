package main

import (
	"github.com/gin-gonic/gin"

	handler "github.com/ismael3s/go-todo/internal/infra/http"
)

func main() {
	r := gin.Default()

	handler.NewRoutesHandler(r).RegisterRoutes()
	r.Run()
}

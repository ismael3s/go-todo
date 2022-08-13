package main

import (
	"log"

	"github.com/gin-gonic/gin"

	handler "github.com/ismael3s/go-todo/internal/infra/http"
	"github.com/ismael3s/go-todo/internal/infra/persistence"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	db := persistence.GetPostgresConnection()

	log.Println("Main", db)
	handler.NewRoutesHandler(r).RegisterRoutes()
	r.Run()
}

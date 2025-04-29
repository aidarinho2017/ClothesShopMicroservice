package main

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/internal/handlers"
	"awesomeProject2/internal/repository"
	"awesomeProject2/internal/service"
	"awesomeProject2/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	database, err := db.NewDB(dsn)
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer database.Close()

	repo := repository.NewItemRepository(database)
	service := service.NewItemService(*repo)
	handler := handlers.NewItemHandler(service)

	r := gin.Default()

	// 👇 Apply CORS middleware
	r.Use(middleware.CORSMiddleware())

	r.POST("/items", handler.CreateItem)
	r.GET("/items/:id", handler.GetItemByID)
	r.PUT("/items/:id", handler.UpdateItem)
	r.DELETE("/items/:id", handler.DeleteItem)

	r.Run(":8080")
}

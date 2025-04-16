package main

import (
	"awesomeProject2/internal"
	"awesomeProject2/internal/db"
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

	repo := internal.NewItemRepository(database)
	service := internal.NewItemService(*repo)
	handler := internal.NewItemHandler(service)

	r := gin.Default()

	r.POST("/items", handler.CreateItem)
	r.GET("/items/:id", handler.GetItemByID)
	r.PUT("/items/:id", handler.UpdateItem)
	r.DELETE("/items/:id", handler.DeleteItem)
	r.Run(":8080")
}

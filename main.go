package main

import (
	"awesomeProject2/internal"
	"awesomeProject2/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/crm?sslmode=disable"

	database, err := db.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	repo := internal.NewPostgresItemRepository(database)
	service := internal.NewItemService(repo)
	handler := internal.NewItemHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/items", func(r chi.Router) {
		r.Post("/", handler.CreateItem)
		r.Get("/shop/{shopId}", handler.GetItems)
		r.Get("/{id}", handler.GetItem)
		r.Put("/", handler.UpdateItem)
		r.Delete("/{id}", handler.DeleteItem)
	})

	log.Println("🚀 Server running at http://localhost:8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}

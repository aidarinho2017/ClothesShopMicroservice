package main

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/internal/handlers"
	"awesomeProject2/internal/repository"
	"awesomeProject2/internal/service"
	"awesomeProject2/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func main() {
	dsn := "postgres://aidarinho:qwerty@localhost:5432/postgres?sslmode=disable"

	database, err := db.NewDB(dsn)
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer database.Close()

	hash := "$2a$10$.eHRgLhZ.TSyOxhfMp/kqud4SflvAt/6QdbTYKCf7X55JGmwHSJNi"
	password := "qwerty"

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("Invalid:", err)
	} else {
		fmt.Println("Success!")
	}

	itemRepo := repository.NewItemRepository(database)
	itemService := service.NewItemService(*itemRepo)
	itemHandler := handlers.NewItemHandler(itemService)

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// 👇 Public auth routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// 👇 Protected item routes
	protected := r.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())

	protected.GET("/profile", func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{"message": "Welcome " + username})
	})

	protected.GET("/items", itemHandler.GetItems)
	protected.POST("/items", itemHandler.CreateItem)
	protected.GET("/items/:id", itemHandler.GetItemByID)
	protected.PUT("/items/:id", itemHandler.UpdateItem)
	protected.DELETE("/items/:id", itemHandler.DeleteItem)

	r.Run(":8080")
}

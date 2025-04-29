package cmd

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/internal/handlers"
	"awesomeProject2/internal/repository"
	"awesomeProject2/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Clothes Shop API
// @version 1.0
// @description This is a sample server for the Clothes Shop API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func Connect() {
	dsn := "postgres://aidarinho:qwerty@localhost:5432/postgres?sslmode=disable"

	database, err := db.NewDB(dsn)
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer database.Close()

	repo := repository.NewItemRepository(database)
	itemService := service.NewItemService(*repo)
	handler := handlers.NewItemHandler(itemService)

	r := gin.Default()

	r.POST("/items", handler.CreateItem)
	r.GET("/items/:id", handler.GetItemByID)
	r.PUT("/items/:id", handler.UpdateItem)
	r.DELETE("/items/:id", handler.DeleteItem)
	r.GET("/items/qr", handler.GetItemByQRCode)        // To identify items by QR code data
	r.GET("/items/:id/qrcode", handler.DownloadQRCode) // To download the QR code as a picture

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

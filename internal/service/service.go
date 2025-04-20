package service

import (
	"awesomeProject2/internal/repository"
	"awesomeProject2/models"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
	"strings"
	"time"
)

type ItemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(ctx context.Context, item *models.Item) error {
	if item.ID == "" {
		item.ID = uuid.New().String()
	}

	// Generate QR code data (the item's unique ID for identification)
	qrCodeData := item.ID

	// Generate the QR code as PNG data
	qrCodeBytes, err := qrcode.Encode(qrCodeData, qrcode.Medium, 256)
	if err != nil {
		log.Printf("❌ Error generating QR code: %v", err)
		return fmt.Errorf("failed to generate QR code: %w", err)
	}

	// Encode the QR code bytes as a base64 string for storage
	item.QRCode = fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(qrCodeBytes))

	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	return s.repo.Create(item)
}

func (s *ItemService) GetItemByID(ctx context.Context, id string) (*models.Item, error) {
	return s.repo.GetByID(id)
}

func (s *ItemService) GetItemByQRCode(ctx context.Context, qrCodeData string) (*models.Item, error) {
	return s.repo.GetByQRCode(qrCodeData)
}

func (s *ItemService) UpdateItem(ctx context.Context, item *models.Item) error {
	item.UpdatedAt = time.Now()
	return s.repo.Update(item)
}

func (s *ItemService) DeleteItem(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}

func (s *ItemService) DownloadQRCode(ctx context.Context, id string, c *gin.Context) {
	item, err := s.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Decode the base64 string to bytes
	dataURL := item.QRCode
	parts := strings.SplitN(dataURL, ",", 2)
	if len(parts) != 2 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid QR code data"})
		return
	}
	base64Data := parts[1]
	qrCodeBytes, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode QR code"})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s_qrcode.png\"", item.ID))
	c.Header("Content-Type", "image/png")
	c.Data(http.StatusOK, "image/png", qrCodeBytes)
}

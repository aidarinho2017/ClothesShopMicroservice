package service

import (
	"awesomeProject2/internal/repository"
	"awesomeProject2/models"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	"log"
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

	// Generate QR code data (e.g., the item's ID or a URL)
	qrCodeData := fmt.Sprintf("item/%s", item.ID) // Example: URL to item details

	// Generate the QR code as PNG data
	qrCodeBytes, err := qrcode.Encode(qrCodeData, qrcode.Medium, 256)
	if err != nil {
		log.Printf("❌ Error generating QR code: %v", err)
		return fmt.Errorf("failed to generate QR code: %w", err)
	}

	// Encode the QR code bytes as a base64 string for easy embedding in JSON
	item.QRCode = fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(qrCodeBytes))

	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	return s.repo.Create(item)
}

func (s *ItemService) GetItemByID(ctx context.Context, id string) (*models.Item, error) {
	return s.repo.GetByID(id)
}

func (s *ItemService) UpdateItem(ctx context.Context, item *models.Item) error {
	return s.repo.Update(item)
}

func (s *ItemService) DeleteItem(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}

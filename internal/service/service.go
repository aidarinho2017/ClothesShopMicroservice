package service

import (
	"awesomeProject2/internal/repository"
	"awesomeProject2/models"
	"context"
	"encoding/base64"
	"fmt"
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
	// No need to manually assign ID

	qrCodeData := fmt.Sprintf("item/%s", item.Name) // or any other unique field

	qrCodeBytes, err := qrcode.Encode(qrCodeData, qrcode.Medium, 256)
	if err != nil {
		log.Printf("❌ Error generating QR code: %v", err)
		return fmt.Errorf("failed to generate QR code: %w", err)
	}

	item.QRCode = fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(qrCodeBytes))
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	return s.repo.Create(item)
}

func (s *ItemService) GetItems(ctx context.Context) ([]*models.Item, error) {
	return s.repo.GetItems()
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

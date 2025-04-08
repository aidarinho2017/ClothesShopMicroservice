package internal

import (
	"awesomeProject2/models"
	"context"
)

type ItemService struct {
	repo ItemRepository
}

func NewItemService(repo ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(ctx context.Context, item *models.Item) error {
	return s.repo.Create(ctx, item)
}

func (s *ItemService) GetItemsForShop(ctx context.Context, shopID int) ([]models.Item, error) {
	return s.repo.GetAll(ctx, shopID)
}

func (s *ItemService) GetItemByID(ctx context.Context, id int) (*models.Item, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ItemService) UpdateItem(ctx context.Context, item models.Item) error {
	return s.repo.Update(ctx, item)
}

func (s *ItemService) DeleteItem(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// Sorted variants 🔽🔼

func (s *ItemService) GetItemsSortedByBrand(ctx context.Context, shopID int) ([]models.Item, error) {
	return s.repo.GetSorted(ctx, shopID, "brand", true)
}

func (s *ItemService) GetItemsSortedByCategory(ctx context.Context, shopID int) ([]models.Item, error) {
	return s.repo.GetSorted(ctx, shopID, "category", true)
}

func (s *ItemService) GetItemsSortedBySize(ctx context.Context, shopID int) ([]models.Item, error) {
	return s.repo.GetSorted(ctx, shopID, "size", true)
}

func (s *ItemService) GetItemsSortedByPrice(ctx context.Context, shopID int, asc bool) ([]models.Item, error) {
	return s.repo.GetSorted(ctx, shopID, "price", asc)
}

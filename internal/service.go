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

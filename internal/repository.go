package internal

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/models"
	"context"
	"fmt"
)

type ItemRepository interface {
	Create(ctx context.Context, item *models.Item) error
	GetAll(ctx context.Context, shopID int) ([]models.Item, error)
	GetByID(ctx context.Context, id int) (*models.Item, error)
	Update(ctx context.Context, item models.Item) error
	Delete(ctx context.Context, id int) error
	GetSorted(ctx context.Context, shopID int, field string, asc bool) ([]models.Item, error)
}

type PostgresItemRepository struct {
	db *db.DB
}

func NewPostgresItemRepository(db *db.DB) *PostgresItemRepository {
	return &PostgresItemRepository{db: db}
}

func (r *PostgresItemRepository) Create(ctx context.Context, item *models.Item) error {
	query := `
		INSERT INTO items (name, brand, category, size, color, price, bought_for, sex, photo, barcode, created_at, updated_at, shop_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW(), $11)
	`
	_, err := r.db.Conn.Exec(ctx, query,
		item.Name,
		item.Brand,
		item.Category,
		item.Size,
		item.Color,
		item.Price,
		item.BoughtFor,
		item.Sex,
		item.Photo,
		item.Barcode,
	)
	return err
}

func (r *PostgresItemRepository) GetAll(ctx context.Context, shopID int) ([]models.Item, error) {
	return r.fetchItems(ctx, fmt.Sprintf(`SELECT * FROM items WHERE shop_id = $1`), shopID)
}

func (r *PostgresItemRepository) GetByID(ctx context.Context, id int) (*models.Item, error) {
	query := `SELECT * FROM items WHERE id = $1`
	var item models.Item
	err := r.db.Conn.QueryRow(ctx, query, id).Scan(
		&item.ID,
		&item.Name,
		&item.Brand,
		&item.Category,
		&item.Size,
		&item.Color,
		&item.Price,
		&item.BoughtFor,
		&item.Sex,
		&item.Photo,
		&item.Barcode,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}
	return &item, nil
}

func (r *PostgresItemRepository) Update(ctx context.Context, item models.Item) error {
	query := `
		UPDATE items SET 
			name = $1, brand = $2, category = $3, size = $4, color = $5, 
			price = $6, bought_for = $7, sex = $8, photo = $9, barcode = $10, 
			updated_at = NOW()
		WHERE id = $11
	`
	_, err := r.db.Conn.Exec(ctx, query,
		item.Name,
		item.Brand,
		item.Category,
		item.Size,
		item.Color,
		item.Price,
		item.BoughtFor,
		item.Sex,
		item.Photo,
		item.Barcode,
		item.ID,
	)
	return err
}

func (r *PostgresItemRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Conn.Exec(ctx, `DELETE FROM items WHERE id = $1`, id)
	return err
}

func (r *PostgresItemRepository) GetSorted(ctx context.Context, shopID int, field string, asc bool) ([]models.Item, error) {
	allowedFields := map[string]bool{
		"brand": true, "category": true, "size": true, "price": true,
	}
	if !allowedFields[field] {
		return nil, fmt.Errorf("invalid sort field")
	}

	order := "ASC"
	if !asc {
		order = "DESC"
	}

	query := fmt.Sprintf(`SELECT * FROM items WHERE shop_id = $1 ORDER BY %s %s`, field, order)
	return r.fetchItems(ctx, query, shopID)
}

func (r *PostgresItemRepository) fetchItems(ctx context.Context, query string, arg int) ([]models.Item, error) {
	rows, err := r.db.Conn.Query(ctx, query, arg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Brand,
			&item.Category,
			&item.Size,
			&item.Color,
			&item.Price,
			&item.BoughtFor,
			&item.Sex,
			&item.Photo,
			&item.Barcode,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

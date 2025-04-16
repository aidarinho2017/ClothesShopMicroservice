package internal

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/models"
	"context"
	"fmt"
	"time"
)

type ItemRepository struct {
	DB *db.DB
}

func NewItemRepository(database *db.DB) *ItemRepository {
	return &ItemRepository{DB: database}
}

func (r *ItemRepository) Create(item *models.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO items (id, name, brand, category, size, color, price, bought_for, sex, photo, qr_code, created_at, updated_at)
    VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
	_, err := r.DB.Conn.Exec(ctx, query,
		item.ID, item.Name, item.Brand, item.Category, item.Size, item.Color,
		item.Price, item.BoughtFor, item.Sex, item.Photo.Photo, item.QRCode, time.Now(), time.Now(),
	)
	if err != nil {
		fmt.Printf("❌ DB Insert Error: %v\n", err)
	}
	return err
}

func (r *ItemRepository) GetByID(id string) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, name, brand, category, size, color, price, bought_for, sex, photo, qr_code, created_at, updated_at FROM items WHERE id=$1`
	row := r.DB.Conn.QueryRow(ctx, query, id)

	var item models.Item
	err := row.Scan(&item.ID, &item.Name, &item.Brand, &item.Category, &item.Size, &item.Color,
		&item.Price, &item.BoughtFor, &item.Sex, &item.Photo.Photo, &item.QRCode, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) Update(item *models.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE items SET name=$1, brand=$2, category=$3, size=$4, color=$5,
       price=$6, bought_for=$7, sex=$8, photo=$9, qr_code=$10, updated_at=$11 WHERE id=$12`
	_, err := r.DB.Conn.Exec(ctx, query,
		item.Name, item.Brand, item.Category, item.Size, item.Color,
		item.Price, item.BoughtFor, item.Sex, item.Photo.Photo, item.QRCode, time.Now(), item.ID,
	)
	return err
}

func (r *ItemRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM items WHERE id=$1`
	_, err := r.DB.Conn.Exec(ctx, query, id)
	return err
}

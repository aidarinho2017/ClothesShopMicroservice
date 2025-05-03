package repository

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/models"
	"context"
	"fmt"
	"log"
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

	query := `INSERT INTO items (name, brand, category, size, color, price, bought_for, sex, photo, qr_code, created_at, updated_at)
              VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
              RETURNING id`

	err := r.DB.Conn.QueryRow(ctx, query,
		item.Name, item.Brand, item.Category, item.Size, item.Color,
		item.Price, item.BoughtFor, item.Sex, item.Photo.Photo, item.QRCode,
		item.CreatedAt, item.UpdatedAt,
	).Scan(&item.ID) // Assign auto-generated ID back to the item

	if err != nil {
		log.Printf("❌ DB Insert Error: %v\n", err)
		return err
	}
	return nil
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

func (r *ItemRepository) GetItems() ([]*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, name, brand, category, size, color, price, bought_for, sex, photo, qr_code, created_at, updated_at FROM items`

	rows, err := r.DB.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Brand, &item.Category, &item.Size, &item.Color,
			&item.Price, &item.BoughtFor, &item.Sex, &item.Photo.Photo, &item.QRCode, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) GetByQRCode(qrCodeData string) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, name, brand, category, size, color, price, bought_for, sex, photo, qr_code, created_at, updated_at FROM items WHERE qr_code LIKE $1`
	row := r.DB.Conn.QueryRow(ctx, query, fmt.Sprintf("data:image/png;base64,%s", qrCodeData))

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

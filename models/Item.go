package models

import "time"

type Item struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Brand     Brand     `json:"brand" db:"brand"`
	Category  Category  `json:"category" db:"category"`
	Size      Size      `json:"size" db:"size"`
	Color     Color     `json:"color" db:"color"`
	Price     int       `json:"price" db:"price"`
	BoughtFor int       `json:"bought_for" db:"bought_for"`
	Sex       string    `json:"sex" db:"sex"`
	Photo     Photo     `json:"photo" db:"photo"`
	Barcode   Barcode   `json:"barcode" db:"barcode"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Barcode struct {
	Code string `json:"code" db:"code"`
}

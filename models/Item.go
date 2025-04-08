package models

import "time"

type Item struct {
	ID        int
	Name      string
	Brand     string
	Category  string
	Size      string
	Color     string
	Price     float64
	BoughtFor float64
	Sex       string
	Photo     string
	Barcode   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Barcode struct {
	Code string `json:"code" db:"code"`
}

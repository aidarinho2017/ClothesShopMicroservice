package models

import "time"

type Item struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Brand     string    `json:"brand" db:"brand"`
	Category  string    `json:"category" db:"category"`
	Size      string    `json:"size" db:"size"`
	Color     string    `json:"color" db:"color"`
	Price     int       `json:"price" db:"price"`
	BoughtFor int       `json:"bought_for" db:"bought_for"`
	Sex       string    `json:"sex" db:"sex"`
	Photo     Photo     `json:"photo" db:"photo"`
	QRCode    string    `json:"qr_code" db:"qr_code"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Photo struct {
	ID    string `json:"id" db:"id"`
	Photo string `json:"photo" db:"photo"` // image URL or base64 string
}

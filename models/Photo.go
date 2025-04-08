package models

type Photo struct {
	ID    string `json:"id" db:"id"`
	Photo string `json:"photo" db:"photo"` // image URL or base64 string
}

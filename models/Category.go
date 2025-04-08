package models

type Category struct {
	Type   string `json:"type" db:"type"`
	Amount int    `json:"amount" db:"amount"` // inventory count
}

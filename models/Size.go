package models

type Size struct {
	Size   string `json:"size" db:"size"`
	Amount int    `json:"amount" db:"amount"`
}

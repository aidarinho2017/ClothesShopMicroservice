package models

type Color struct {
	Color  string `json:"color" db:"color"`
	Amount int    `json:"amount" db:"amount"`
}

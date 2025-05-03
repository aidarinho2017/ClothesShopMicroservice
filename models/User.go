package models

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Age            int    `json:"age"`
	Identification string `json:"identification"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
}

package models

type User struct {
	Email   string `json:"email"`
	PassKey string `json:"pass"`
}

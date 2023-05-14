package models

type Member struct {
	ID         int    `json:"user_id"`
	Mail       string `json:"email"`
	Permission int    `json:"permission"`
}

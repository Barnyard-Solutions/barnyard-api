package models

type User struct {
	ID      int    `json:"ID"`
	Email   string `json:"email"`
	PassKey string `json:"pass"`
}

type SafeUser struct {
	ID    int    `json:"ID"`
	Email string `json:"email"`
}

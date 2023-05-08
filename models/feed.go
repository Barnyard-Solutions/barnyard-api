package models

type Feed struct {
	ID         int    `json:"ID"`
	Name       string `json:"name"`
	Permission int    `json:"permission"`
}

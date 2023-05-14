package models

type Feed struct {
	ID         int    `json:"ID"`
	Name       string `json:"name"`
	Permission int    `json:"permission"`
}

type FeedSub struct {
	ID         int    `json:"ID"`
	Name       string `json:"name"`
	Permission int    `json:"permission"`
	Subscribed bool   `json:"subscribed"`
}

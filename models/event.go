package models

type Event struct {
	ID     int    `json:"ID"`
	Name1  string `json:"name1"`
	Name2  string `json:"name2"`
	Date   string `json:"date"`
	FeedID int    `json:"feed_id"`
}

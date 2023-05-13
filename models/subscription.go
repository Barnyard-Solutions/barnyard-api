package models

type Subscription struct {
	UserID   int    `json:"user_id"`
	FeedID   int    `json:"feed_id"`
	EndPoint string `json:"end_point"`
}

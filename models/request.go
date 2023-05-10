package models

type CreateFeedRequest struct {
	Token    string `json:"token"`
	FeedName string `json:"feed_name"`
}

type DeleteFeedRequest struct {
	Token  string `json:"token"`
	FeedID int    `json:"feed_id"`
}

type CreateEventRequest struct {
	Token  string `json:"token"`
	FeedID int    `json:"feed_id"`
	Name1  string `json:"name1"`
	Name2  string `json:"name2"`
	Date   string `json:"date"`
}

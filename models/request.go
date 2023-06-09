package models

type CreateFeedRequest struct {
	FeedName string `json:"feed_name"`
}

type CreateEventRequest struct {
	Name1 string `json:"name1"`
	Name2 string `json:"name2"`
	Date  string `json:"date"`
}

type CreateMilestoneRequest struct {
	Name  string `json:"name"`
	Date  int    `json:"date"`
	Color string `json:"color"`
}

type SubscriptionRequest struct {
	Subscription string `json:"subscription"`
}

type MemberRequest struct {
	Mail       string `json:"email"`
	Permission int    `json:"permission"`
}

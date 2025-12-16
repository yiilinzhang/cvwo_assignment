package models


type Topic struct {
	ID   int    `json:"topic_id"`
	Title string `json:"title"`
	UserId string `json:"user_id"`
}

package topics

type CreateTopicInput struct {
	Title  string `json:"title"`
	UserID int    `json:"user_id"`
}

type DeleteTopicInput struct {
	TopicID int `json:"topic_id"`
	UserID  int `json:"user_id"`
}

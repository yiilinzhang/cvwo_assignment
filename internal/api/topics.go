package api
type CreateTopicInput struct {
	Title string `json:"title"`;
	UserId int `json:"user_id"`;	
}

type DeleteTopicInput struct {
	TopicId string `json:"topic_id"`;
	UserId int `json:"user_id"`;	
}

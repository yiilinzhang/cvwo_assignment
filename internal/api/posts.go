package api

type CreatePostInput struct {
	TopicId int `json:"topic_id"`;
	Title string `json:"title"`;
	Content string `json:"content"`;
	UserId int `json:"user_id"`;	
}
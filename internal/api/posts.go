package api

type CreatePostInput struct {
	TopicId int `json:"topic_id"`;
	Title string `json:"title"`;
	Content string `json:"content"`;
	UserId int `json:"user_id"`;	
}

type DeletePostInput struct {
	PostId int `json:"post_id"`;
	UserId int `json:"user_id"`;	
}
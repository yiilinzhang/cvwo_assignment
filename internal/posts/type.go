package posts

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

type UpdatePostInput struct {
	PostId int `json:"post_id"`;
	Content string `json:"content"`;
	UserId int `json:"user_id"`;	 
}

type QueryPostResponse struct {
	ID      int    `json:"post_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
	TopicTitle string `json:"topic_title" `
}

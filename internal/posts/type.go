package posts

type CreatePostInput struct {
	TopicID int    `json:"topic_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

type DeletePostInput struct {
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
}

type UpdatePostInput struct {
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

type QueryPostResponse struct {
	ID         int    `json:"post_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	UserID     string `json:"user_id"`
	TopicTitle string `json:"topic_title" `
}

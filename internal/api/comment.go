package api

type CreateCommentInput struct {
	Content string `json:"content"`
	UserId  int `json:"user_id"`
	PostId int `json:"post_id" `
}

type DeleteCommentInput struct {
	CommentId int `json:"comment_id"`;
	UserId int `json:"user_id"`;	
}

type UpdateCommentInput struct {
	CommentId int `json:"comment_id"`;
	Content string `json:"content"`;
	UserId int `json:"user_id"`;	 
}
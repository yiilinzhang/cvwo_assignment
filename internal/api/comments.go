package api

type CreateCommentInput struct {
	Content string `json:"content"`
	UserId  string `json:"user_id"`
	PostId string `json:"post_id" `
}
package api

type CreateCommentInput struct {
	Content string `json:"content"`
	UserId  int `json:"user_id"`
	PostId string `json:"post_id" `
}
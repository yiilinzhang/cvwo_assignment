package models

type Comment struct {
	ID      int    `json:"comment_id"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
	PostId string `json:"post_id" `
}

//What i want returned, user.username, post.content, post.title, comment.content, comment.id 

//TODO check if i shoudl move this to api
type QueryCommentResponse struct {
	ID      int    `json:"comment_id"`
	Content string `json:"content"`
	UserName string `json:"name"`
}

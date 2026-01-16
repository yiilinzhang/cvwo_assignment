package models

type Comment struct {
	ID      int    `json:"comment_id"`
	Content string `json:"content"`
	UserId  int `json:"user_id"`
	PostId int `json:"post_id" `
}

//What i want returned, user.username, post.content, post.title, comment.content, comment.id 

//TODO check if i shoudl move this to api
type QueryCommentResponse struct {
	ID      int    `json:"comment_id"`
	Content string `json:"content"`
	UserName string `json:"name"`
	UserId  int `json:"user_id"`
}

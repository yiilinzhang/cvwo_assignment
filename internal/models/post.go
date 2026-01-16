package models

//TODO add required userID later when login implemented
//TODO look into implementing struc validaiton maybe with the payground validate
type Post struct {
	ID      int    `json:"post_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int `json:"user_id"`
	TopicId int `json:"topic_id" `
}

type QueryPostResponse struct {
	ID      int    `json:"post_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
	TopicTitle string `json:"topic_title" `
}

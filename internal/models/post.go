package models

//TODO add required userID later when login implemented
//TODO look into implementing struc validaiton maybe with the payground validate
type Post struct {
	ID      int    `json:"post_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
	TopicId string `json:"topic_id" `
}

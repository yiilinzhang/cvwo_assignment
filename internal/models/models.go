package models

// Represents a fourum user
type User struct {
	ID           int    `json:"userid"`
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
}

// Represents a discussion topic (one topic -> many posts under the topic)
type Topic struct {
	ID     int    `json:"topic_id"`
	Title  string `json:"title"`
	UserId int    `json:"user_id"`
}

// What i want returned, user.username, post.content, post.title, comment.content, comment.id
// TODO add required userID later when login implemented
// TODO look into implementing struc validaiton maybe with the payground validate
// Represents a post under a topic
type Post struct {
	ID      int    `json:"post_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	TopicId int    `json:"topic_id" `
}

// Represents a comment under a post
type Comment struct {
	ID      int    `json:"comment_id"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	PostId  int    `json:"post_id" `
}

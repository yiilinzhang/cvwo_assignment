package comments

type CreateCommentInput struct {
	Content  string `json:"content"`
	UserID   int    `json:"user_id"`
	PostID   int    `json:"post_id,omitempty"`
	ParentID *int   `json:"parent_comment_id,omitempty"`
}

type DeleteCommentInput struct {
	CommentID int `json:"comment_id"`
	UserID    int `json:"user_id"`
}

type UpdateCommentInput struct {
	CommentID int    `json:"comment_id"`
	Content   string `json:"content"`
	UserID    int    `json:"user_id"`
}

type QueryCommentResponse struct {
	ID       int    `json:"comment_id"`
	Content  string `json:"content"`
	UserName string `json:"name"`
	UserID   int    `json:"user_id"`
	ParentID *int   `json:"parent_comment_id,omitempty"`
}

type CommentParent struct {
	ID       int             `json:"comment_id"`
	Content  string          `json:"content"`
	UserName string          `json:"name"`
	UserID   int             `json:"user_id"`
	ParentID *int            `json:"parent_comment_id,omitempty"`
	Children []*CommentParent `json:"children"`
}
